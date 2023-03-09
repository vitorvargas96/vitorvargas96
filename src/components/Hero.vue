<template>
    <div class="h-screen w-screen flex flex-col space-y-8 items-center justify-center text-center overflow-hidden">
        <img :src="Avatar" class="relative rounded-full h-32 w-32 mx-auto object-cover">
        <div>
            <h2 class="text-sm uppercase text-gray-500 pb-2 tracking-[15px]">
                Full Stack Developer
            </h2>

            <h1 class="text-5xl lg:text-6xl font-semibold px-10">
                Vitor Vargas
            </h1>
        </div>

        <div>
            <a href="#aboutme">
                <button class="navBtn">About</button>
            </a>
            <a href="#experience">
                <button class="navBtn">Experience</button>
            </a>
            <a href="#skills">
                <button class="navBtn">Skills</button>
            </a>
            <a href="#projects">
                <button class="navBtn">Projects</button>
            </a>
        </div>
    </div>
</template>

<script setup lang="ts">
import Avatar from '../assets/avatar.jpg'
</script>