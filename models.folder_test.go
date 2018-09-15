package main

import "testing"

func TestGetUsersFolders(t *testing.T) {
    flist := getUsersFolders()

    // Check length of folders
    if len(flist) != len(folderList) {
        t.Fail()
    }

    // Make sure each member is identical
    for i, v := range flist {
        if v.Owner != folderList[i].Owner ||
        v.FolderName != folderList[i].FolderName ||
        v.Created != folderList[i].Created ||
        v.NumElements != folderList[i].NumElements ||
        v.Completed != folderList[i].Completed ||
        v.Downloaded != folderList[i].Downloaded {
            
        t.Fail()
        break
        }
    }
}