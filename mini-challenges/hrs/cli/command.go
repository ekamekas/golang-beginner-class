package cli

import "fmt"
import "strings"
import "hrs/db"

func Parse(database *db.Db, args []string) {
    // guard
    if (len(args) < 1) {
        throw();
    } 

    if("help" == args[0]) {
        helpCommand();
    } else if("search" == args[0]) {
        searchCommand(database, args[1:]);
    } else if("pretty" == args[0]) {
        prettyCommand(database);
    } else if("add" == args[0]) {
        panic("not implemented);
    } else {
        throw();
    }
}

func throw() {
        panic("Argument tidak valid. Silahkan gunakan `help` untuk mengetahui komando");
}

func helpCommand() {
    fmt.Println("[HELP]\nKomando yang dapat digunakan:\nsearch name [name] -> mencari berdasarkan nama\nsearch id [id] -> mencari berdasarkan id\nadd -> menambahkan data ke database\npretty [n] -> menampilkan seluruh data");
}

func searchCommand(database *db.Db, args []string) {
    // guard
    if (len(args) < 2) {
        throw();
    } 
    
    if("name" == args[0]) {
        person, _ := database.FindByName(strings.Join(args[1:], " "));

        fmt.Printf("%+v", person);
    } else if("id" == args[0] && len(args) != 2) {
        throw();
    } else if("id" == args[0]) {
        person, _ := database.FindById(args[1]);

        fmt.Printf("%+v", person);
    } else {
        throw();
    }
}

func prettyCommand(database *db.Db) {
    fmt.Printf("Nama\t\t\tAlamat\t\tPekerjaan\t\tAlasan Bergabung\n\n");
    for _, person := range database.IndexByName {
        fmt.Printf("%s\t\t%s\t\t%s\t\t%s\n", person.Name, person.Address, person.Job, person.Reason);
    }
}
