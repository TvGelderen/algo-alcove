<template>
    <header class="bg-secondary flex w-full">
        <div>
            <nav class="text-2xl m-4">
                <RouterLink to="/" class="hover:text-theme">Home</RouterLink>
            </nav>
        </div>
    </header>

    <div class="w-[95%] max-w-[1440px] mx-auto my-4">
    <RouterView />
    </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>
