package repositories

import(
    "DHTWeather/models"
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type DhtReadingRepository struct {
    db *sql.DB
}

func NewDhtReadingRepository() DhtReadingRepository {
    return DhtReadingRepository {
        db: createCon(),
    }
}

func (rp *DhtReadingRepository) Close() {
    rp.db.Close()
}

func (rp *DhtReadingRepository) Save(r models.Reading) error {
    err := InsertReading(rp.db, r)
    return err
}

/*Create mysql connection*/
func createCon() *sql.DB {
    db, err := sql.Open("mysql", "newuser:password@tcp(localhost:3306)/DHT")
    
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("db is connected")
    
    return db
}

func InsertReading(db *sql.DB, r models.Reading) error {
    // perform a db.Query insert
    // TODO: Convert to Stored Procedures
    insertQuery := fmt.Sprintf("INSERT INTO readings(createdate, nodeid, temperature, humidity) VALUES ( '%s', %d, %f, %f )", r.ReadingDate, r.NodeId, r.Temperature, r.Humidity)
    fmt.Println(insertQuery)
    
    insert, err := db.Query(insertQuery)
    
    fmt.Println("Inserting Data...");
    
    // if there is an error inserting, handle it
    if err != nil {
        return err
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()
    
    return nil
}

func FindAllReadings(db *sql.DB) ([]models.Reading, error) {
    selectQuery := fmt.Sprintf("SELECT id, createdate, nodeid, temperature, humidity FROM readings")
    
    sel, err := db.Query(selectQuery)
    
    if err != nil {
        return []models.Reading{}, err
    }
    readings := []models.Reading{};
    
    for sel.Next() {
        var r models.Reading
        
        // for each row, scan the result into our object
        err = sel.Scan(&r.ID, &r.ReadingDate, &r.NodeId, &r.Temperature, &r.Humidity)  //Mutating, so use pointer
        if err != nil {
            return []models.Reading{}, err
        }
        readings = append(readings, r)
    }
    
    defer sel.Close()
    
    return readings, nil
}
