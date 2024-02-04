<template>
    <header class="bg-secondary w-full h-[42px] md:h-[60px] relative flex justify-between items-center px-2"
        data-open="false">
        <div class="flex items-center">
            <div class="text-2xl md:text-3xl">
                AlgoAlcove
            </div>
            <nav class="hidden md:flex text-xl md:text-2xl m-4">
                <RouterLink to="/" class="hover:text-theme">Home</RouterLink>
            </nav>
        </div>
        <div class="md:hidden" id="menu-icon">
            <span></span>
            <span></span>
            <span></span>
        </div>
    </header>

    <div class="w-[clamp(360px,97%,1500px)] mx-auto my-4">
        <RouterView />
    </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { RouterLink, RouterView } from 'vue-router'

onMounted(() => {
    const header = document.querySelector('header');
    const menuIcon = document.getElementById('menu-icon');

    if (menuIcon && header) {
        const resizeListener = () => {
            if (window.innerWidth > 768) {
                header.dataset.open = "false";
                window.removeEventListener('resize', resizeListener);
            }
        }

        menuIcon.addEventListener('click', () => {
            if (header.dataset.open == "true") {
                header.dataset.open = "false";
                window.removeEventListener('resize', resizeListener);
            } else {
                header.dataset.open = "true";
                window.addEventListener('resize', resizeListener);
            }
        });
    }
});
</script>

<style scoped>
#menu-icon {
    width: 32px;
    height: 32px;
}

#menu-icon>span {
    display: block;
    background: #ffffff;
    border-radius: 10px;
    height: 3px;
    margin: 6px 0;
    transition: .4s cubic-bezier(0.68, -0.6, 0.32, 1.6);
}

#menu-icon>span:nth-of-type(1) {
    width: 50%;
}

#menu-icon>span:nth-of-type(2) {
    width: 100%;
}

#menu-icon>span:nth-of-type(3) {
    width: 75%;
}

header[data-open="true"] #menu-icon>span:nth-of-type(1) {
    transform-origin: bottom;
    transform: rotateZ(45deg) translate(4px, 0.5px);
}

header[data-open="true"] #menu-icon>span:nth-of-type(2) {
    transform-origin: top;
    transform: rotateZ(-45deg);
}

header[data-open="true"] #menu-icon>span:nth-of-type(3) {
    width: 50%;
    transform-origin: bottom;
    transform: translate(13.5px, -4px) rotateZ(45deg);
}
</style>
