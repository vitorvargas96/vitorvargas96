<template>
    <header class="sticky top-0 p-5 flex items-start justify-between max-w-7xl mx-auto z-20 xl:items-center">
        <div class="flex flex-row items-center space-x-8">
            <a href="https://github.com/vitorvargasdev" target="_blank">
                <div v-html="GithubIcon" class="fill-gray-300 w-7 hover:fill-white" />
            </a>

            <a href="https://t.me/vitorvargasdev" target="_blank">
                <div v-html="TelegramIcon" class="fill-gray-300 w-7 hover:fill-white" />
            </a>

            <a href="" target="_blank">
                <div v-html="CoffeeIcon" class="fill-gray-300 w-7 hover:fill-white" />
            </a>
        </div>
        <div class="flex flex-row items-center mr-3">
            <a href="https://www.linkedin.com/in/vitorvargasdev/" target="_blank" class="py-2 px-5 border rounded-md border-gray-400 hover:bg-neutral-900">
                LinkedIn
            </a>
        </div>
    </header>
</template>

<script setup lang="ts">
import GithubIcon from '../assets/github.svg?raw'
import TelegramIcon from '../assets/telegram.svg?raw'
import CoffeeIcon from '../assets/coffee.svg?raw'
</script>