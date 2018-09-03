package models

import (
    "time"
)

type Folder struct {
    modelImpl
    FolderName   string
    Created      time.Date
    NumElements  int
}
