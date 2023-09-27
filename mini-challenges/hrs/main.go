package main

import "os"
import "hrs/db"
import "hrs/cli"

func main() {
    args := os.Args[1:];

    dbInstance := &db.Db{
        IndexById: make(map[string]*db.PersonEntity),
        IndexByName: make(map[string]*db.PersonEntity),
    }

    dataInitialization(dbInstance);

    cli.Parse(dbInstance, args);
}

func dataInitialization(database *db.Db) {
    persons := []db.PersonEntity{
        { Id: "1", Name: "Mas Eka Setiawan", Job: "na", Address: "na", Reason: "na" },
        { Id: "2", Name: "Leslie E Smith", Job: "Business Owner", Address: "Ohio", Reason: "na" },
        { Id: "3", Name: "Susan S Harjo", Job: "Metal Worker", Address: "Washington", Reason: "na" },
        { Id: "4", Name: "Robert H Jones", Job: "Cost Estimator", Address: "Bandung", Reason: "na" },
    }

    for _, person := range persons {
        database.Insert(person);
    }
}
