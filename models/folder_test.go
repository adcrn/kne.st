package models

import "testing"

func TestGetUsersFolders(t *testing.T) {
	flist := GetUsersFolders(3)

	// Check length of folders
	if len(flist) != len(Folders) {
		t.Fail()
	}

	// Make sure each member is identical
	for i, v := range flist {
		if v.OwnerID != Folders[i].OwnerID ||
			v.FolderName != Folders[i].FolderName ||
			v.Created != Folders[i].Created ||
			v.NumElements != Folders[i].NumElements ||
			v.Completed != Folders[i].Completed ||
			v.Downloaded != Folders[i].Downloaded {

			t.Fail()
			break
		}
	}
}
