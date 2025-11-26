<template>
  <div class="min-h-screen flex flex-col items-center justify-between p-5 relative z-10">
    <!-- Header -->
    <header class="text-center mb-10 pt-5">
      <h1 class="text-4xl md:text-5xl font-bold text-isan-pink-600 drop-shadow-lg mb-3">
        🌸 เกรียนเวฟ App 🌸
      </h1>
      <p class="text-lg text-isan-pink-500 font-medium">
        แอปพื้นบ้านอีสาน สำหรับเกรียนเวฟโดยเฉพาะ
      </p>
    </header>

    <!-- Main Content -->
    <div class="flex flex-col items-center gap-8">
      <!-- Main Button -->
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

      <!-- Coming Soon Section -->
      <div class="bg-white/60 backdrop-blur-sm rounded-2xl p-4 max-w-sm text-center border-2 border-isan-pink-300 shadow-lg mt-4">
        <div class="flex items-center justify-center gap-2 mb-2">
          <span class="text-2xl">🚀</span>
          <h3 class="text-lg font-bold text-isan-pink-600">Coming Soon!</h3>
          <span class="text-2xl">🌟</span>
        </div>
        <p class="text-sm text-isan-pink-500 leading-relaxed">
          เร็วๆ นี้! <span class="font-bold text-isan-orange">Wave Pop E3</span> 🐱✨
          <br/>
          คล้าย Pop Cat แต่เป็นเวอร์ชันอีสาน!
          <br/>
          มาร่วมกดสร้างสถิติโลกกันเถอะ! 🏆
          <br/>
          <span class="text-xs text-isan-pink-400 mt-1 block">ทุกคนร่วมกดทำคะแนนได้ทั่วโลก!</span>
        </p>
      </div>
    </div>

    <!-- Footer -->
    <footer class="mt-8 pb-5 text-center">
      <!-- Social Links -->
      <div class="flex items-center justify-center gap-6 mb-4">
        <!-- GitHub -->
        <a 
          href="https://github.com/khonE3" 
          target="_blank" 
          rel="noopener noreferrer"
          class="group flex flex-col items-center gap-1 transition-transform hover:scale-110"
        >
          <div class="w-12 h-12 rounded-full bg-gradient-to-br from-gray-700 to-gray-900 
                      flex items-center justify-center shadow-lg border-2 border-isan-gold
                      group-hover:shadow-[0_0_20px_rgba(255,215,0,0.5)]">
            <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
            </svg>
          </div>
          <span class="text-xs font-medium text-isan-pink-600">khonE3</span>
        </a>

        <!-- Instagram -->
        <a 
          href="https://www.instagram.com/gotjitag/" 
          target="_blank" 
          rel="noopener noreferrer"
          class="group flex flex-col items-center gap-1 transition-transform hover:scale-110"
        >
          <div class="w-12 h-12 rounded-full bg-gradient-to-br from-purple-500 via-pink-500 to-orange-400 
                      flex items-center justify-center shadow-lg border-2 border-isan-gold
                      group-hover:shadow-[0_0_20px_rgba(255,105,180,0.5)]">
            <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"/>
            </svg>
          </div>
          <span class="text-xs font-medium text-isan-pink-600">gotjitag</span>
        </a>
      </div>

      <!-- Copyright -->
      <div class="text-isan-pink-500 text-sm">
        <p class="font-medium">Made with 💖 in หนองบัวลำภู</p>
        <p class="text-xs text-isan-pink-400 mt-1">© 2025 เกรียนเวฟ App | E3 Team</p>
      </div>
    </footer>

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

// 🎵 Isan Sound Effects using Web Audio API
let audioContext: AudioContext | null = null

const initAudio = () => {
  if (!audioContext) {
    audioContext = new (window.AudioContext || (window as any).webkitAudioContext)()
  }
  return audioContext
}

// เสียงพิณอีสาน (Isan Phin sound)
const playIsanSound = () => {
  const ctx = initAudio()
  if (!ctx) return

  const now = ctx.currentTime

  // สุ่มเสียงแบบอีสาน
  const sounds = [
    () => playPhinSound(ctx, now),      // เสียงพิณ
    () => playKhaenSound(ctx, now),     // เสียงแคน
    () => playPongLangSound(ctx, now),  // เสียงโปงลาง
  ]
  
  sounds[Math.floor(Math.random() * sounds.length)]()
}

// เสียงพิณ (Phin - Isan lute)
const playPhinSound = (ctx: AudioContext, now: number) => {
  const notes = [329.63, 392, 440, 523.25] // E4, G4, A4, C5
  const note = notes[Math.floor(Math.random() * notes.length)]
  
  const osc = ctx.createOscillator()
  const gain = ctx.createGain()
  
  osc.type = 'triangle'
  osc.frequency.setValueAtTime(note, now)
  osc.frequency.exponentialRampToValueAtTime(note * 1.02, now + 0.1)
  
  gain.gain.setValueAtTime(0.4, now)
  gain.gain.exponentialRampToValueAtTime(0.01, now + 0.5)
  
  osc.connect(gain)
  gain.connect(ctx.destination)
  
  osc.start(now)
  osc.stop(now + 0.5)
}

// เสียงแคน (Khaen - bamboo mouth organ)
const playKhaenSound = (ctx: AudioContext, now: number) => {
  const baseFreq = 220 + Math.random() * 110 // A3-A4
  
  for (let i = 0; i < 3; i++) {
    const osc = ctx.createOscillator()
    const gain = ctx.createGain()
    
    osc.type = 'sine'
    osc.frequency.setValueAtTime(baseFreq * (i + 1), now)
    
    gain.gain.setValueAtTime(0.2 / (i + 1), now)
    gain.gain.exponentialRampToValueAtTime(0.01, now + 0.4)
    
    osc.connect(gain)
    gain.connect(ctx.destination)
    
    osc.start(now)
    osc.stop(now + 0.4)
  }
}

// เสียงโปงลาง (Pong Lang - wooden xylophone)
const playPongLangSound = (ctx: AudioContext, now: number) => {
  const notes = [523.25, 587.33, 659.25, 783.99] // C5, D5, E5, G5
  
  notes.slice(0, 2 + Math.floor(Math.random() * 2)).forEach((note, i) => {
    const osc = ctx.createOscillator()
    const gain = ctx.createGain()
    
    osc.type = 'square'
    osc.frequency.setValueAtTime(note, now + i * 0.08)
    
    gain.gain.setValueAtTime(0, now + i * 0.08)
    gain.gain.linearRampToValueAtTime(0.25, now + i * 0.08 + 0.02)
    gain.gain.exponentialRampToValueAtTime(0.01, now + i * 0.08 + 0.2)
    
    osc.connect(gain)
    gain.connect(ctx.destination)
    
    osc.start(now + i * 0.08)
    osc.stop(now + i * 0.08 + 0.25)
  })
}

const teaseWave = async () => {
  // 🎵 Play Isan sound!
  playIsanSound()
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
