package models

import (
    "errors"
    )

type Reading struct {
    ID int
    NodeId int
    ReadingDate string
    Temperature float32
    Humidity float32
}

func NewReading(id int, nodeId int, readingDate string, temperature float32, humidity float32) (Reading, error) {
    if nodeId < 0 {
        return Reading{}, errors.New("nodeId should be greater than 0")
    }
    
    return Reading {
        ID: id,
        NodeId: nodeId,
        ReadingDate: readingDate,
        Temperature: temperature,
        Humidity: humidity,
        }, nil
}
