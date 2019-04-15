/*
	Copyright © 2019 The Goca.io team

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package goca

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"

	jsoncolor "github.com/nwidger/jsoncolor"
	"github.com/timshannon/bolthold"
	bson "go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	PS             ProjectStore
	CurrentProject Project
)

type ProjectStore struct {
	DS *bolthold.Store
}

type Project struct {
	ID      bson.ObjectID `boltholdKey:"ID"`
	Name    string
	Created time.Time
	Scans   []bson.ObjectID
}

type Scan struct {
	ID       bson.ObjectID `boltholdKey:"ID"`
	Output   Output
	Started  time.Time
	Finished time.Time
}

// Opens a Project Store
func OpenProjectStore() (ProjectStore, error) {
	var ps ProjectStore

	store, err := bolthold.Open("./goca.db", 0666, nil)
	if err != nil {
		return ps, err
	}

	ps.DS = store

	return ps, nil
}

// Creates a new project
func (ps ProjectStore) NewProject(name string) (Project, error) {
	id := bson.NewObjectID()
	project := Project{
		ID:      id,
		Name:    name,
		Created: time.Now(),
	}

	err := ps.DS.Insert(bson.NewObjectID(), project)
	if err != nil {
		return project, err
	}

	return project, nil
}

// Obtain a project
func (ps ProjectStore) GetProject(name string) (Project, error) {
	var projects []Project

	err := ps.DS.Find(&projects, bolthold.Where("Name").Eq(name))
	if err != nil {
		return Project{}, err
	}

	if len(projects) < 1 {
		return Project{}, errors.New("Project not found")
	}

	return projects[0], nil
}

// Adds an scan to an existing project
func (ps ProjectStore) AddScanToProject(name string, scan bson.ObjectID) error {
	project, err := ps.GetProject(name)
	if err != nil {
		return err
	}

	newProject := Project{
		ID:      project.ID,
		Name:    project.Name,
		Created: project.Created,
		Scans:   append(project.Scans, scan),
	}

	err = ps.DS.Update(project.ID, newProject)
	if err != nil {
		return err
	}

	return nil
}

// Deletes an existing project
func (ps ProjectStore) DeleteProject(name string) error {
	project, err := ps.GetProject(name)
	if err != nil {
		return err
	}

	err = ps.DS.Delete(project.ID, project)
	if err != nil {
		return err
	}

	return nil
}

// Saves an scan associated to a project
func (ps ProjectStore) SaveScan(project string, output *Output) error {
	id := bson.NewObjectID()
	scan := Scan{
		ID:       id,
		Finished: time.Now(),
		Output:   *output,
	}

	err := ps.DS.Insert(id, scan)
	if err != nil {
		return err
	}

	err = ps.AddScanToProject(project, id)
	if err != nil {
		return err
	}

	return nil
}

// Gets an scan
func (ps ProjectStore) GetScan(id bson.ObjectID) (Scan, error) {
	var scan Scan

	err := ps.DS.Get(id, &scan)
	if err != nil {
		return Scan{}, err
	}

	return scan, nil
}

// Prints a project
func (ps ProjectStore) PrintProject(name string) error {
	var scans []Scan

	project, err := ps.GetProject(name)
	if err != nil {
		return err
	}

	for _, scanid := range project.Scans {
		scan, err := ps.GetScan(scanid)
		if err != nil {
			return err
		}
		scans = append(scans, scan)
	}

	if runtime.GOOS == "windows" {
		data, err := json.MarshalIndent(scans, "", "\t")
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Println(string(data))
		}
	} else {
		data, err := jsoncolor.MarshalIndent(scans, "", "\t")
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Println(string(data))
		}
	}

	return nil
}
