<template>
    <main>
        <div :hx-get="endpoint" hx-trigger="load"></div>
    </main>
</template>

<script setup lang="ts">
const endpoint = import.meta.env.VITE_API_ENDPOINT;
</script>
