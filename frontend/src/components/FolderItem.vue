<template>
  <div class="card">
    <header class="card-header">
      <p class="card-header-title">
        <a v-bind:href="folder.path">{{ folder.folder_name }}</a>
      </p>
    </header>
    <div class="card-content">
      <div class="content">
        Uploaded: {{ folder.time }}<br>
        <i>{{ folder.num_elements }} items</i><br>
        <!--Completed: {{ folder.completed }}<br>
        Downloaded: {{ folder.downloaded }}<br>-->
      </div>
    </div>
    <footer class="card-footer">
      <a href="#" class="card-footer-item">Download</a>
      <a href="#" class="card-footer-item">Delete</a>
    </footer>
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
