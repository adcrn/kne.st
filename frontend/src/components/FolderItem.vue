<template>
  <div>
    <a v-bind:href="folder.path">{{ folder.folder_name }}</a>
    Uploaded: {{ folder.time }}<br>
    <i>{{ folder.num_elements }} items</i><br>
    Completed: {{ folder.completed }}<br>
    Downloaded: {{ folder.downloaded }}<br>
  </div>
</template>

<script>
export default {
  props: {
    folder: {
      type: Object,
      required: true
    }
  }
};
</script>
