<template>
  <div class="folders flex mb-4">
    <div class="columns" v-if="folder_list.length">
      <div class="column w-full bg-white h-12" v-for="folder in folder_list" v-bind:key="folder.id">
        <FolderItem
          v-bind:folder="folder"
        />
      </div>
    </div>
    <p v-else>
      There's nothing here!<br>
      <button class="button is-link">Upload a folder</button>
    </p>
  </div>
</template>

<script>
const fetch_folders_url = "http://localhost:9000/fetch/:id";
let nextfolderid = 1;

import axios from "axios";
import FolderItem from "../components/FolderItem.vue";

export default {
  name: "folders",
  components: {
    FolderItem
  },
  data() {
    return {
      error: "",
      folder_list: [
        {
          id: nextfolderid++,
          owner_id: 1,
          folder_name: "terns",
          time: "2018-01-20",
          path: "./storage/1/terns",
          num_elements: 12,
          completed: false,
          downloaded: false
        },
        {
          id: nextfolderid++,
          owner_id: 1,
          folder_name: "gulls",
          time: "2018-02-20",
          path: "./storage/1/gulls",
          num_elements: 17,
          completed: false,
          downloaded: false
        },
        {
          id: nextfolderid++,
          owner_id: 1,
          folder_name: "test3",
          time: "2018-01-20",
          path: "./storage/1/terns",
          num_elements: 12,
          completed: false,
          downloaded: false
        },
        {
          id: nextfolderid++,
          owner_id: 1,
          folder_name: "test4",
          time: "2018-02-20",
          path: "./storage/1/gulls",
          num_elements: 17,
          completed: false,
          downloaded: false
        }
      ]
    };
  },
  computed: {
    noFolders() {
      return this.folder_list.length === 0;
    }
  },
  mounted() {
    axios
      .get(fetch_folders_url)
      .then(response => {
        this.folder_list = response.data;
      })
      .catch(e => {
        this.errors.push(e);
      });
  }
};
</script>

<style scoped lang="scss">
p {
  padding: 2rem;
}

.folders {
  padding: 3rem;
}
</style>
