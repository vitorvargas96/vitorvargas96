<template>
    <div class="h-screen w-screen flex flex-col relative text-center md:text-left md:flex-row max-w-7xl px-10 justify-evenly mx-auto items-center overflow-hidden">
        <img :src="Avatar" class="-mb-20 flex justify-start relative rounded-full h-56 w-56 object-cover flex-shrink-0 md:rounded-lg md:w-64 md:h-95 xl:w-[500px] xl:h-[500px]">
        
        <div class="space-y-10 px-0 md:px-10">
            <h4 class="text-4xl font-semibold">
                Hello world!
            </h4>
            <p class="text-base">
                I'm a passionate, self-taught software developer. My insatiable curiosity and desire to constantly learn lead me to seek out new knowledge and explore innovative technologies.
                <br />
                <br />
                Overall, I'm a highly motivated and dedicated developer who takes pride in delivering exceptional results. I'm excited to continue learning, growing, and contributing to the world of technology in meaningful ways.
            </p>
        </div>
    </div>
</template>

<script setup lang="ts">
import Avatar from '../assets/avatar.jpg'
</script>