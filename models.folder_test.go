package main

import "testing"

func TestGetUsersFolders(t *testing.T) {
	flist := getUsersFolders(3)

	// Check length of folders
	if len(flist) != len(folders) {
		t.Fail()
	}

	// Make sure each member is identical
	for i, v := range flist {
		if v.OwnerID != folders[i].OwnerID ||
			v.FolderName != folders[i].FolderName ||
			v.Created != folders[i].Created ||
			v.NumElements != folders[i].NumElements ||
			v.Completed != folders[i].Completed ||
			v.Downloaded != folders[i].Downloaded {

			t.Fail()
			break
		}
	}
}
