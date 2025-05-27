<template>
  <div :class="['register-container', { 'dark-mode': darkMode }]">
    <div class="register-card">
      <div class="theme-switcher" @click="toggleTheme">
        <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24">
          <path v-if="darkMode" fill="currentColor" d="M20 8.69V4h-4.69L12 .69 8.69 4H4v4.69L.69 12 4 15.31V20h4.69L12 23.31 15.31 20H20v-4.69L23.31 12 20 8.69zM12 18c-3.31 0-6-2.69-6-6s2.69-6 6-6 6 2.69 6 6-2.69 6-6 6z"/>
          <path v-else fill="currentColor" d="M20 8.69V4h-4.69L12 .69 8.69 4H4v4.69L.69 12 4 15.31V20h4.69L12 23.31 15.31 20H20v-4.69L23.31 12 20 8.69zM12 18c-3.31 0-6-2.69-6-6s2.69-6 6-6 6 2.69 6 6-2.69 6-6 6zm0-10c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4z"/>
        </svg>
      </div>

      <div class="register-header">
        <h2>Create Account</h2>
        <p>Fill in your details to get started</p>
      </div>

      <form @submit.prevent="register" class="register-form">
        <div class="input-group">
          <label for="name">Full Name</label>
          <input
            id="name"
            v-model="name"
            type="text"
            placeholder="Enter your full name"
            required
            class="input-field"
            @input="validateForm"
          />
          <span class="input-icon email-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
              <circle cx="12" cy="7" r="4"></circle>
            </svg>
          </span>
        </div>

        <div class="input-group">
          <label for="email">Email Address</label>
          <input
            id="email"
            v-model="email"
            type="email"
            placeholder="Enter your email"
            required
            class="input-field"
            @input="validateEmail"
            :class="{ 'error': emailError }"
          />
          <span class="input-icon email-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path>
              <polyline points="22,6 12,13 2,6"></polyline>
            </svg>
          </span>
          <div v-if="emailError" class="input-error">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <line x1="12" y1="8" x2="12" y2="12"></line>
              <line x1="12" y1="16" x2="12.01" y2="16"></line>
            </svg>
            {{ emailError }}
          </div>
        </div>

        <div class="input-group">
          <label for="password">Password</label>
          <div class="password-wrapper">
            <input
              id="password"
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="Create a password"
              required
              class="input-field"
              @input="checkPasswordStrength"
            />
            <span class="input-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
                <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
              </svg>
            </span>
            <span class="toggle-password" @click="showPassword = !showPassword">
              <svg v-if="showPassword" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path>
                <line x1="1" y1="1" x2="23" y2="23"></line>
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                <circle cx="12" cy="12" r="3"></circle>
              </svg>
            </span>
          </div>
          <div class="password-strength">
            <div
              class="strength-bar"
              :class="{
                'weak': passwordStrength === 'weak',
                'medium': passwordStrength === 'medium',
                'strong': passwordStrength === 'strong'
              }"
              :style="{ width: passwordStrengthPercent + '%' }"
            ></div>
            <span class="strength-text">{{ passwordStrengthText }}</span>
          </div>
        </div>

        <button
          type="submit"
          class="register-button"
          :disabled="!formValid || loading"
          :class="{ 'disabled': !formValid }"
        >
          <span v-if="!loading">Create Account</span>
          <span v-else class="loader"></span>
        </button>

        <div v-if="error" class="error-message">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="10"></circle>
            <line x1="12" y1="8" x2="12" y2="12"></line>
            <line x1="12" y1="16" x2="12.01" y2="16"></line>
          </svg>
          {{ error }}
        </div>
      </form>

      <div class="register-footer">
        <p>Already have an account? <router-link to="/login">Sign in</router-link></p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import api from '@/axios'
import { useRouter } from 'vue-router'

const name = ref('')
const email = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)
const showPassword = ref(false)
const darkMode = ref(true)
const passwordEntropy = ref(0)
const formValid = ref(false)

const router = useRouter()

function getCookie(name: string): string | null {
  const value = `; ${document.cookie}`
  const parts = value.split(`; ${name}=`)
  if (parts.length === 2) return parts.pop()?.split(';').shift() || null
  return null
}

function setCookie(name: string, value: string, days = 365) {
  const date = new Date()
  date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000))
  document.cookie = `${name}=${value};expires=${date.toUTCString()};path=/;SameSite=Lax`
}

onMounted(() => {
  const themeCookie = getCookie('theme')
  if (themeCookie) {
    darkMode.value = themeCookie === 'dark'
  } else {
    darkMode.value = true
  }
  updateTheme()
})

onMounted(() => {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme) {
    darkMode.value = savedTheme === 'dark'
    updateTheme()
  } else {
    darkMode.value = window.matchMedia('(prefers-color-scheme: dark)').matches
    updateTheme()
  }
})

const passwordStrength = computed(() => {
  if (passwordEntropy.value >= 65) return 'strong'
  if (passwordEntropy.value >= 50) return 'medium'
  return 'weak'
})

const passwordStrengthPercent = computed(() => {
  return Math.min(Math.max(passwordEntropy.value, 0), 100)
})

const passwordStrengthText = computed(() => {
  switch (passwordStrength.value) {
    case 'weak': return 'Weak password (entropy < 50)'
    case 'medium': return `Medium strength (entropy ${passwordEntropy.value.toFixed(1)})`
    case 'strong': return `Strong password (entropy ${passwordEntropy.value.toFixed(1)})`
    default: return ''
  }
})

function validateForm() {
  formValid.value = name.value.trim() !== '' &&
    email.value.trim() !== '' &&
    password.value.trim() !== '' &&
    passwordStrength.value !== 'weak'
}

function calculatePasswordEntropy(password: string): number {
  if (!password) return 0;

  const uniqueChars = new Set(password).size;
  const length = password.length;

  let entropy = 0;
  for (let i = 0; i < uniqueChars; i++) {
    entropy += Math.log2(length);
  }
  for (let i = 0; i < length - uniqueChars; i++) {
    entropy += Math.log2(uniqueChars || 1);
  }

  return entropy;
}

function checkPasswordStrength() {
  passwordEntropy.value = calculatePasswordEntropy(password.value);
  validateForm();
}

async function register() {
  try {
    error.value = ''
    loading.value = true

    if (!formValid.value) {
      throw new Error('Please fill all fields correctly')
    }

    const response = await api.post('/v1/register/', {
      name: name.value,
      email: email.value,
      password: password.value
    })

    document.cookie = `auth_token=${response.data.token}; path=/; SameSite=Lax; max-age=86400`
    router.push('/')

  } catch (err) {
    error.value = err.response?.data?.message || err.message || 'Registration failed'
  } finally {
    loading.value = false
  }
}

function toggleTheme() {
  darkMode.value = !darkMode.value
  updateTheme()
  setCookie('theme', darkMode.value ? 'dark' : 'light')
}

function updateTheme() {
  if (darkMode.value) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
}

const emailValid = computed(() => {
  const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return re.test(email.value)
})

const emailError = ref('')

function validateEmail() {
  if (!email.value) {
    emailError.value = ''
    return
  }

  if (!emailValid.value) {
    emailError.value = 'Please enter a valid email address'
  } else {
    emailError.value = ''
  }

  validateForm()
}

</script>

<style scoped>
.register-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  background-color: #f5f3ff;
  transition: all 0.3s ease;
  overflow-y: auto;
}

.dark-mode.register-container {
  background-color: #1e1b4b;
}

.register-card {
  width: 100%;
  max-width: 500px;
  padding: 3.5rem;
  border-radius: 16px;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
  background-color: white;
  position: relative;
  z-index: 1;
  transition: all 0.3s ease;
  margin: auto;
}

.theme-switcher {
  position: absolute;
  top: 1.5rem;
  right: 1.5rem;
  cursor: pointer;
  color: #7e22ce;
  transition: all 0.3s ease;
}

.theme-switcher:hover {
  transform: scale(1.1);
  color: #6d28d9;
}

.register-header {
  text-align: center;
  margin-bottom: 2.5rem;
}

.register-header h2 {
  color: #7e22ce;
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  letter-spacing: -0.025em;
}

.register-header p {
  color: #6b7280;
  font-size: 1rem;
  line-height: 1.5;
  font-weight: 400;
  max-width: 320px;
  margin: 0 auto;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
}

.register-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.input-group {
  position: relative;
}

.input-group label {
  display: block;
  margin-bottom: 0.75rem;
  color: #28006d;
  font-size: 0.9375rem;
  font-weight: 500;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  letter-spacing: 0.01em;
}

.input-field {
  width: 100%;
  padding: 1rem 1.25rem 1rem 3.25rem;
  border: 1px solid #e5e7eb;
  border-radius: 10px;
  font-size: 1rem;
  transition: all 0.3s ease;
  background-color: #f9fafb;
  height: 3.25rem;
  box-sizing: border-box;
  color: #111827;
}

.input-field:focus {
  outline: none;
  border-color: #8b5cf6;
  box-shadow: 0 0 0 3px rgba(139, 92, 246, 0.2);
}

.input-icon {
  position: absolute;
  left: 1rem;
  top: 50%;
  transform: translateY(-50%);
  color: #8b5cf6;
  height: 1.5rem;
  width: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.input-icon.email-icon {
  transform: none;
  top: 55%;
}

.password-wrapper {
  position: relative;
}

.toggle-password {
  position: absolute;
  right: 1rem;
  top: 50%;
  transform: translateY(-50%);
  cursor: pointer;
  color: #8b5cf6;
  height: 1.5rem;
  width: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.password-strength {
  margin-top: 0.5rem;
  height: 6px;
  background-color: #e5e7eb;
  border-radius: 3px;
  overflow: hidden;
  position: relative;
}

.strength-bar {
  height: 100%;
  position: absolute;
  left: 0;
  top: 0;
  transition: width 0.3s ease, background-color 0.3s ease;
}

.strength-bar.weak {
  background-color: #ef4444;
}

.strength-bar.medium {
  background-color: #f59e0b;
}

.strength-bar.strong {
  background-color: #10b981;
}

.strength-text {
  display: block;
  margin-top: 0.25rem;
  font-size: 0.75rem;
  color: #6b7280;
}

.register-button {
  background-color: #7e22ce;
  color: white;
  border: none;
  padding: 1.125rem;
  border-radius: 10px;
  font-size: 1.0625rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  height: 3.5rem;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 0.5rem;
}

.register-button:hover:not(:disabled) {
  background-color: #6d28d9;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(126, 34, 206, 0.2);
}

.register-button:disabled {
  background-color: #d1d5db;
  cursor: not-allowed;
  transform: none !important;
  box-shadow: none !important;
}

.error-message {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #ef4444;
  font-size: 0.9375rem;
  padding: 0.875rem;
  background-color: #fee2e2;
  border-radius: 8px;
  margin-top: 0.5rem;
}

.register-footer {
  margin-top: 2rem;
  text-align: center;
  font-size: 0.9375rem;
  color: #6b7280;
}

.register-footer a {
  color: #7e22ce;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.2s;
}

.register-footer a:hover {
  color: #6d28d9;
  text-decoration: underline;
}

.loader {
  width: 1.5rem;
  height: 1.5rem;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  border-top-color: white;
  animation: spin 1s ease-in-out infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.dark-mode .register-card {
  background-color: #312e81;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.3);
}

.dark-mode .register-header h2 {
  color: #a78bfa;
}

.dark-mode .register-header p {
  color: #c7d2fe;
}

.dark-mode .input-group label {
  color: #e0e7ff;
}

.dark-mode .input-field {
  background-color: #4338ca;
  border-color: #4c1d95;
  color: #e9d5ff;
}

.dark-mode .input-field:focus {
  border-color: #a78bfa;
  box-shadow: 0 0 0 3px rgba(167, 139, 250, 0.3);
}

.dark-mode .input-icon,
.dark-mode .toggle-password {
  color: #a78bfa;
}

.dark-mode .password-strength {
  background-color: #374151;
}

.dark-mode .strength-text {
  color: #9ca3af;
}

.dark-mode .register-button {
  background-color: #7c3aed;
}

.dark-mode .register-button:hover:not(:disabled) {
  background-color: #8b5cf6;
}

.dark-mode .register-button:disabled {
  background-color: #28006d;
}

.dark-mode .error-message {
  color: #fca5a5;
  background-color: #7f1d1d;
}

.dark-mode .register-footer {
  color: #c7d2fe;
}

.dark-mode .register-footer a {
  color: #a78bfa;
}

@media (max-width: 640px) {
  .register-container {
    align-items: flex-start;
    padding: 3rem 1rem 1rem;
  }

  .register-card {
    padding: 2.5rem 1.5rem;
    max-width: 100%;
  }

  .register-header h2 {
    font-size: 1.75rem;
  }

  .input-field {
    padding: 0.875rem 1rem 0.875rem 3rem;
    height: 3rem;
  }

  .register-button {
    padding: 1rem;
    height: 3.25rem;
  }
}
</style>
