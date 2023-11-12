package time

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2"
	lru "github.com/hashicorp/golang-lru"
	"github.com/life4/genesis/slices"
	enterpriseos "github.com/sigmasee/sigmasee/shared/enterprise/os"
	"github.com/sigmasee/sigmasee/shared/enterprise/tuples"
)

// Many systems use /usr/share/zoneinfo, Solaris 2 has
// /usr/share/lib/zoneinfo, IRIX 6 has /usr/lib/locale/TZ,
// NixOS has /etc/zoneinfo.
var platformZoneSources = []string{
	"/usr/share/zoneinfo/",
	"/usr/share/lib/zoneinfo/",
	"/usr/lib/locale/TZ/",
	"/etc/zoneinfo",
}

type TimeHelper interface {
	Max() time.Time
	StartOfDay(timezone ...string) carbon.Carbon
	EndOfDay(dateTime *carbon.Carbon, timezone ...string) carbon.Carbon
	EndOfDayFromStandardDate(dateTime time.Time) time.Time
	ToShortDateWithoutYear(dateTime carbon.Carbon) string
	RemoveAllPartsAfterSeconds(dateTime carbon.Carbon) carbon.Carbon
	GetAllTimezoneLocations() ([]string, error)
	GetNowMinusOneDay() time.Time
}

type timeHelper struct {
	osHelper               enterpriseos.OsHelper
	timezoneLocationsCache *lru.Cache
}

func NewTimeHelper(osHelper enterpriseos.OsHelper) (TimeHelper, error) {
	timezoneLocationsCache, err := lru.New(1)
	if err != nil {
		return nil, err
	}

	return &timeHelper{
		osHelper:               osHelper,
		timezoneLocationsCache: timezoneLocationsCache,
	}, nil
}

func (s *timeHelper) Max() time.Time {
	return time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC)
}

func (s *timeHelper) StartOfDay(timezone ...string) carbon.Carbon {
	return s.RemoveAllPartsAfterSeconds(carbon.Now(timezone...).SetTimeNano(0, 0, 0, 0))
}

func (s *timeHelper) EndOfDay(dateTime *carbon.Carbon, timezone ...string) carbon.Carbon {
	var startOfDay carbon.Carbon

	if dateTime == nil {
		startOfDay = s.StartOfDay(timezone...)
	} else {
		startOfDay = *dateTime
	}

	return s.RemoveAllPartsAfterSeconds(startOfDay.SetTimeNano(23, 59, 59, 0))
}

func (s *timeHelper) EndOfDayFromStandardDate(dateTime time.Time) time.Time {
	return s.RemoveAllPartsAfterSeconds(carbon.CreateFromStdTime(dateTime).SetTimeNano(23, 59, 59, 0)).ToStdTime()
}

func (s *timeHelper) ToShortDateWithoutYear(dateTime carbon.Carbon) string {
	return dateTime.Format("l jS F")
}

func (s *timeHelper) RemoveAllPartsAfterSeconds(dateTime carbon.Carbon) carbon.Carbon {
	return dateTime.
		SetMillisecond(0).
		SetMicrosecond(0).
		SetNanosecond(0)
}

func (s *timeHelper) GetAllTimezoneLocations() ([]string, error) {
	const key string = "all"

	if cachedResult, ok := s.timezoneLocationsCache.Get(key); ok {
		return cachedResult.([]string), nil
	}

	result := slices.Map(platformZoneSources, func(item string) tuples.ValueErrorTuple[[]string] {
		timezones, err := s.readFile(item, "")
		if err != nil {
			return tuples.ValueErrorTuple[[]string]{
				Value: nil,
				Error: err,
			}
		}

		return tuples.ValueErrorTuple[[]string]{
			Value: timezones,
			Error: nil,
		}
	})

	if err := tuples.ReduceErrors(result); err != nil {
		return nil, err
	}

	timezones := tuples.GetValues(result)

	timezoneLocations := slices.Reduce(timezones, []string{}, func(items []string, acc []string) []string {
		return append(acc, items...)
	})

	_ = s.timezoneLocationsCache.Add(key, timezoneLocations)

	return timezoneLocations, nil
}

func (s *timeHelper) GetNowMinusOneDay() time.Time {
	return time.Now().UTC().Add(-24 * time.Hour)
}

func (s *timeHelper) readFile(zoneDir, path string) ([]string, error) {
	finalPath := filepath.Join(zoneDir, path)

	if !s.osHelper.DirExist(finalPath) {
		return []string{}, nil
	}

	files, err := os.ReadDir(finalPath)
	if err != nil {
		return nil, err
	}

	timezones := []string{}

	for _, file := range files {
		if file.Name() != strings.ToUpper(file.Name()[:1])+file.Name()[1:] {
			continue
		}

		if file.IsDir() {
			nestedTimezones, err := s.readFile(zoneDir, filepath.Join(path, file.Name()))
			if err != nil {
				return nil, err
			}

			timezones = append(timezones, nestedTimezones...)
		} else {
			if path != "" {
				timezones = append(timezones, fmt.Sprintf("%s/%s", path, file.Name()))
			}
		}
	}

	return timezones, nil
}
