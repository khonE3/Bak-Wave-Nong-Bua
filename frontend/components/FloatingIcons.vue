<template>
  <div class="fixed inset-0 pointer-events-none z-0 overflow-hidden">
    <span 
      v-for="(icon, index) in floatingIcons" 
      :key="index"
      class="absolute text-3xl opacity-60"
      :class="icon.animation"
      :style="icon.style"
    >
      {{ icon.emoji }}
    </span>
  </div>
</template>

<script setup lang="ts">
const floatingIcons = [
  { emoji: '🌾', style: { top: '10%', left: '5%' }, animation: 'animate-float' },
  { emoji: '🐃', style: { top: '20%', right: '10%' }, animation: 'animate-float-slow' },
  { emoji: '🌸', style: { top: '40%', left: '8%' }, animation: 'animate-float-fast' },
  { emoji: '🪭', style: { top: '60%', right: '15%' }, animation: 'animate-float' },
  { emoji: '🥁', style: { top: '80%', left: '12%' }, animation: 'animate-float-slow' },
  { emoji: '🌺', style: { top: '15%', left: '75%' }, animation: 'animate-float-fast' },
  { emoji: '🦎', style: { top: '50%', right: '8%' }, animation: 'animate-float' },
  { emoji: '🍚', style: { top: '75%', left: '65%' }, animation: 'animate-float-slow' },
  { emoji: '🎵', style: { top: '30%', left: '45%' }, animation: 'animate-float-fast' },
  { emoji: '💃', style: { top: '85%', right: '25%' }, animation: 'animate-float' },
]
</script>
