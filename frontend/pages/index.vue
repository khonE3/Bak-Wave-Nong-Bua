<template>
  <div class="min-h-screen flex flex-col items-center justify-center p-5 relative z-10">
    <!-- Header -->
    <header class="text-center mb-10">
      <h1 class="text-4xl md:text-5xl font-bold text-isan-pink-600 drop-shadow-lg mb-3">
        🌸 เกรียนเวฟ App 🌸
      </h1>
      <p class="text-lg text-isan-pink-500 font-medium">
        แอปพื้นบ้านอีสาน สำหรับเกรียนเวฟโดยเฉพาะ
      </p>
    </header>

    <!-- Main Button -->
    <div class="flex flex-col items-center gap-8">
      <button
        @click="teaseWave"
        :class="{ 'animate-shake': isShaking }"
        class="relative w-52 h-52 rounded-full bg-gradient-to-br from-isan-pink-400 to-isan-pink-600 
               border-4 border-isan-gold shadow-2xl 
               flex flex-col items-center justify-center
               transition-all duration-300 ease-out
               hover:scale-110 hover:shadow-[0_0_40px_rgba(255,105,180,0.7)]
               active:scale-95
               animate-pulse-glow
               overflow-hidden"
      >
        <!-- Shimmer effect -->
        <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent 
                    animate-shimmer bg-[length:200%_100%]"></div>
        
        <span class="text-5xl mb-2 relative z-10">👆</span>
        <span class="text-xl font-bold text-white drop-shadow-lg relative z-10">กดเกรียนเวฟ!</span>
      </button>

      <!-- Counter -->
      <div class="bg-white/80 backdrop-blur-sm px-6 py-3 rounded-full border-3 border-isan-gold shadow-lg">
        <p class="text-xl font-bold text-isan-pink-600">
          เกรียนไปแล้ว: <span class="text-isan-orange">{{ teaseCount }}</span> ครั้ง
        </p>
      </div>
    </div>

    <!-- Message Popup -->
    <Transition name="popup">
      <div
        v-if="showMessage"
        @click="hideMessage"
        class="fixed inset-0 flex items-center justify-center z-50 bg-black/20"
      >
        <div class="bg-gradient-to-br from-isan-gold to-isan-orange p-8 rounded-3xl 
                    border-4 border-isan-pink-400 shadow-2xl
                    animate-pop-in text-center max-w-sm mx-4">
          <p class="text-2xl font-bold text-red-800 mb-4">{{ currentMessage }}</p>
          <span class="text-6xl block animate-bounce-slow">{{ currentEmoji }}</span>
        </div>
      </div>
    </Transition>

    <!-- Confetti Container -->
    <div id="confetti-container" class="fixed inset-0 pointer-events-none z-40"></div>
  </div>
</template>

<script setup lang="ts">
// Page metadata
definePageMeta({
  title: 'เกรียนเวฟ App',
  layout: 'default'
})

const teaseCount = ref(0)
const showMessage = ref(false)
const isShaking = ref(false)
const currentMessage = ref('')
const currentEmoji = ref('')

const messages = [
  'หนองบัวลำภูอีสานมากๆ!',
  'เวฟ! อีสานบ้านเฮา!',
  'สิบักหลาย เวฟเอ้ย!',
  'หนองบัวลำภู แดนอีสาน!',
  'เวฟจังซี่! อีสานสุดๆ!',
  'ลำซิ่งหนองบัวลำภู!',
  'เวฟเอ๋ย อีสานบ่?',
  'หนองบัวลำภู ดีใจหลาย!',
  'เวฟบักคัก อีสานแท้!',
  'หนองบัวลำภู ม่วนซื่น!'
]

const emojis = ['🌾', '🐃', '🌸', '🪭', '🥁', '🌺', '🎵', '🎉', '💃', '🕺', '🦎', '🍚']

const teaseWave = async () => {
  // Shake animation
  isShaking.value = true
  setTimeout(() => {
    isShaking.value = false
  }, 500)

  // Random message
  currentMessage.value = messages[Math.floor(Math.random() * messages.length)]
  currentEmoji.value = emojis[Math.floor(Math.random() * emojis.length)]

  // Show message popup
  showMessage.value = true

  // Create confetti
  createConfetti()

  // Increment counter
  teaseCount.value++

  // Call backend API
  try {
    const config = useRuntimeConfig()
    await $fetch(`${config.public.apiBase}/api/tease`, {
      method: 'POST',
      body: { name: 'Wave' }
    })
  } catch (error) {
    console.log('Backend not connected, running in offline mode')
  }

  // Auto hide message
  setTimeout(() => {
    hideMessage()
  }, 2000)
}

const hideMessage = () => {
  showMessage.value = false
}

const createConfetti = () => {
  const container = document.getElementById('confetti-container')
  if (!container) return

  const colors = ['#FF69B4', '#FFB6C1', '#FFD700', '#FF8C00', '#FFC0CB', '#DB7093']

  for (let i = 0; i < 30; i++) {
    const confetti = document.createElement('div')
    confetti.className = 'fixed w-4 h-4 rounded-full animate-confetti-fall'
    confetti.style.left = Math.random() * 100 + 'vw'
    confetti.style.backgroundColor = colors[Math.floor(Math.random() * colors.length)]
    confetti.style.animationDuration = (Math.random() * 2 + 1) + 's'
    confetti.style.animationDelay = Math.random() * 0.5 + 's'
    container.appendChild(confetti)

    // Remove confetti after animation
    setTimeout(() => {
      confetti.remove()
    }, 3500)
  }
}
</script>

<style scoped>
.popup-enter-active {
  animation: pop-in 0.5s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.popup-leave-active {
  animation: pop-out 0.3s ease-out;
}

@keyframes pop-out {
  from {
    transform: scale(1);
    opacity: 1;
  }
  to {
    transform: scale(0.5);
    opacity: 0;
  }
}

.border-3 {
  border-width: 3px;
}
</style>
