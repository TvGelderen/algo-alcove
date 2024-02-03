<template>
  <header class="bg-secondary flex w-full">
    <div>
      <nav class="text-2xl m-4">
        <RouterLink to="/" class="hover:text-theme">Home</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<style scoped>
</style>
