package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sachaos/toggl/lib"
)

var c *Cache

type Cache struct {
	Filename string
	Content  CacheContent
}

type CacheContent struct {
	CurrentTimeEntry toggl.TimeEntry `json:"current_time_entry"`
}

func New(filename string) *Cache {
	c = new(Cache)
	c.Filename = filename
	return c
}

func Init() { c.Init() }
func (c *Cache) Init() error {
	if err := c.Read(); err != nil {
		if err = c.Write(); err != nil {
			return err
		}
	}
	return nil
}

func Read() { c.Read() }
func (c *Cache) Read() error {
	jsonString, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonString, &c.Content); err != nil {
		return err
	}
	return nil
}

func Write() { c.Write() }
func (c *Cache) Write() error {
	buf, err := json.MarshalIndent(c.Content, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(c.Filename, buf, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func SetCurrentTimeEntry(timeEntry toggl.TimeEntry) { c.SetCurrentTimeEntry(timeEntry) }
func (c *Cache) SetCurrentTimeEntry(timeEntry toggl.TimeEntry) {
	c.Content.CurrentTimeEntry = timeEntry
}

func GetContent() CacheContent { return c.GetContent() }
func (c *Cache) GetContent() CacheContent {
	return c.Content
}
