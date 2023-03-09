<template>
  <div id="main" class="scroll-smooth	bg-neutral-800 h-screen w-screen text-white snap-y snap-mandatory overflow-scroll z-0 scrollbar scrollbar-track-gray-400/20 scrollbar-thumb-red-500">
    <Header />

    <section id="hero" class="snap-always snap-center w-screen h-screen">
      <Hero />
    </section>

    <section id="aboutme" class="snap-always snap-center w-screen h-screen">
      <Aboutme />
    </section>

    <section id="experience" class="snap-always snap-center w-screen h-screen">
      <Experience />
    </section>

    <section id="skills" class="snap-always snap-center w-screen h-screen">
      <Skills />
    </section>

    <section id="projects" class="snap-always snap-center w-screen h-screen">
      <Projects />
    </section>
  </div>
</template>

<script setup lang="ts">
import Header from './components/Header.vue'
import Hero from './components/Hero.vue'
import Aboutme from './components/Aboutme.vue'
import Experience from './components/Experience.vue'
import Projects from './components/Projects.vue'
import Skills from './components/Skills.vue'
</script>
