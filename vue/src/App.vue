<template>
  <router-view />
  <GlobalPlayerBar />
  <router-view name="modal" />
</template>

<script setup>
import GlobalPlayerBar from "./components/GlobalPlayerBar.vue";
</script>
