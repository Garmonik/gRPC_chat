<template>
  <div :class="['login-container', { 'dark-mode': darkMode }]">
    <div class="login-card">
      <div class="theme-switcher" @click="toggleTheme">
        <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24">
          <path v-if="darkMode" fill="currentColor" d="M20 8.69V4h-4.69L12 .69 8.69 4H4v4.69L.69 12 4 15.31V20h4.69L12 23.31 15.31 20H20v-4.69L23.31 12 20 8.69zM12 18c-3.31 0-6-2.69-6-6s2.69-6 6-6 6 2.69 6 6-2.69 6-6 6z"/>
          <path v-else fill="currentColor" d="M20 8.69V4h-4.69L12 .69 8.69 4H4v4.69L.69 12 4 15.31V20h4.69L12 23.31 15.31 20H20v-4.69L23.31 12 20 8.69zM12 18c-3.31 0-6-2.69-6-6s2.69-6 6-6 6 2.69 6 6-2.69 6-6 6zm0-10c-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4-1.79-4-4-4z"/>
        </svg>
      </div>

      <div class="login-header">
        <h2>Welcome Back</h2>
        <p>Please enter your credentials to access your account</p>
      </div>

      <form @submit.prevent="login" class="login-form">
        <div class="input-group">
          <label for="username">Email Address</label>
          <input
            id="username"
            v-model="username"
            type="email"
            placeholder="Enter your email address"
            required
            class="input-field"
          />
          <span class="input-icon email-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path>
              <polyline points="22,6 12,13 2,6"></polyline>
            </svg>
          </span>
        </div>

        <div class="input-group">
          <label for="password">Password</label>
          <div class="password-wrapper">
            <input
              id="password"
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="Enter your password"
              required
              class="input-field"
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
        </div>

        <div class="form-options">
          <div class="remember-me">
            <input type="checkbox" id="remember" v-model="rememberMe" />
            <label for="remember">Remember me</label>
          </div>
          <router-link to="/forgot-password" class="forgot-password">Forgot password?</router-link>
        </div>

        <button type="submit" class="login-button">
          <span v-if="!loading">Sign In</span>
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

      <div class="login-footer">
        <p>Don't have an account? <router-link to="/register">Create account</router-link></p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, watchEffect } from 'vue'
import api from '@/axios'
import { useRouter } from 'vue-router'

const username = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)
const showPassword = ref(false)
const rememberMe = ref(false)
const darkMode = ref(true)
const router = useRouter()

watchEffect(() => {
  document.title = 'Login'
})

onMounted(() => {
  const savedTheme = document.cookie.split('; ').find(row => row.startsWith('theme='))
  if (savedTheme) {
    darkMode.value = savedTheme.split('=')[1] === 'dark'
    updateTheme()
  } else {
    darkMode.value = window.matchMedia('(prefers-color-scheme: dark)').matches
    updateTheme()
  }

  const savedCredentials = document.cookie.split('; ').find(row => row.startsWith('remembered_credentials='))
  if (savedCredentials) {
    try {
      const credentials = JSON.parse(decodeURIComponent(savedCredentials.split('=')[1]))
      username.value = credentials.email
      password.value = credentials.password
      rememberMe.value = true
    } catch (e) {
      console.error('Error parsing saved credentials', e)
    }
  }
})

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

  const savedCredentials = getCookie('remembered_credentials')
  if (savedCredentials) {
    try {
      const credentials = JSON.parse(decodeURIComponent(savedCredentials))
      username.value = credentials.email
      password.value = credentials.password
      rememberMe.value = true
    } catch (e) {
      console.error('Error parsing saved credentials', e)
    }
  }
})

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

async function login() {
  try {
    error.value = ''
    loading.value = true

    const response = await api.post('/v1/login/', {
      email: username.value,
      password: password.value,
    })

    document.cookie = `auth_token=${response.data.token}; path=/; SameSite=Lax; max-age=86400`

    if (rememberMe.value) {
      const credentials = { email: username.value, password: password.value }
      const expires = new Date()
      expires.setFullYear(expires.getFullYear() + 1)
      document.cookie = `remembered_credentials=${encodeURIComponent(JSON.stringify(credentials))}; expires=${expires.toUTCString()}; path=/`
    } else {
      document.cookie = 'remembered_credentials=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/'
    }

    router.push('/')

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (err) {
    error.value = 'Invalid email or password'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
}

#app {
  min-height: 100vh;
}

.login-container {
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

.dark-mode.login-container {
  background-color: #1e1b4b;
}

.login-card {
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

.login-header {
  text-align: center;
  margin-bottom: 2.5rem;
}

/* Обновленные стили для текстов */
.login-header h2 {
  color: #7e22ce;
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  letter-spacing: -0.025em;
}

.login-header p {
  color: #6b7280;
  font-size: 1rem;
  line-height: 1.5;
  font-weight: 400;
  max-width: 320px;
  margin: 0 auto;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
}

.input-group label {
  display: block;
  margin-bottom: 0.75rem;
  color: #4b5563;
  font-size: 0.9375rem;
  font-weight: 500;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  letter-spacing: 0.01em;
}

.input-field::placeholder {
  color: #9ca3af;
  font-weight: 400;
  letter-spacing: 0.01em;
}

.remember-me label {
  color: #4b5563;
  font-size: 0.9375rem;
  font-weight: 500;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  letter-spacing: 0.01em;
}

.forgot-password {
  color: #7e22ce;
  font-size: 0.9375rem;
  font-weight: 500;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  letter-spacing: 0.01em;
}

.login-button {
  font-size: 1.0625rem;
  font-weight: 600;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  letter-spacing: 0.02em;
}

.login-footer p {
  color: #6b7280;
  font-size: 0.9375rem;
  font-weight: 400;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
}

.login-footer a {
  color: #7e22ce;
  font-weight: 500;
}

/* Dark mode текстовые стили */
.dark-mode .login-header h2 {
  color: #a78bfa;
}

.dark-mode .login-header p {
  color: #c7d2fe;
}

.dark-mode .input-group label {
  color: #e0e7ff;
}

.dark-mode .input-field::placeholder {
  color: #a5b4fc;
}

.dark-mode .remember-me label {
  color: #e0e7ff;
}

.dark-mode .login-footer p {
  color: #c7d2fe;
}

.login-header h2,
.login-header p,
.input-group label,
.input-field::placeholder,
.remember-me label,
.forgot-password,
.login-button,
.login-footer p {
  transition: all 0.3s ease;
}

.login-form {
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
  color: #4b5563;
  font-size: 0.9375rem;
  font-weight: 500;
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

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: -0.5rem;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9375rem;
  color: #4b5563;
}

.remember-me input {
  accent-color: #8b5cf6;
  width: 1.125rem;
  height: 1.125rem;
  margin: 0;
  cursor: pointer;
}

.forgot-password {
  color: #7e22ce;
  font-size: 0.9375rem;
  font-weight: 500;
  text-decoration: none;
  transition: color 0.2s;
}

.forgot-password:hover {
  color: #6d28d9;
  text-decoration: underline;
}

.login-button {
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

.login-button:hover {
  background-color: #6d28d9;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(126, 34, 206, 0.2);
}

.login-button:active {
  transform: translateY(0);
  box-shadow: none;
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

.login-footer {
  margin-top: 2rem;
  text-align: center;
  font-size: 0.9375rem;
  color: #6b7280;
}

.login-footer a {
  color: #7e22ce;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.2s;
}

.login-footer a:hover {
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

.login-container.dark-mode {
  background-color: #1e1b4b;
}

.dark-mode .login-card {
  background-color: #312e81;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.3);
}

.dark-mode .theme-switcher {
  color: #7e22ce;
}

.dark-mode .theme-switcher:hover {
  color: #8b5cf6;
}

.dark-mode .login-header h2 {
  color: #a78bfa;
}

.dark-mode .login-header p {
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

.dark-mode .remember-me {
  color: #e0e7ff;
}

.dark-mode .forgot-password {
  color: #a78bfa;
}

.dark-mode .forgot-password:hover {
  color: #8b5cf6;
}

.dark-mode .login-button {
  background-color: #7c3aed;
}

.dark-mode .login-button:hover {
  background-color: #8b5cf6;
}

.dark-mode .error-message {
  color: #fca5a5;
  background-color: #7f1d1d;
}

.dark-mode .login-footer {
  color: #c7d2fe;
}

.dark-mode .login-footer a {
  color: #a78bfa;
}

@media (max-width: 640px) {
  .login-container {
    align-items: flex-start;
    padding: 3rem 1rem 1rem;
  }

  .login-card {
    padding: 2.5rem 1.5rem;
    max-width: 100%;
  }

  .login-header h2 {
    font-size: 1.75rem;
  }

  .input-field {
    padding: 0.875rem 1rem 0.875rem 3rem;
    height: 3rem;
  }

  .login-button {
    padding: 1rem;
    height: 3.25rem;
  }
}

@media (min-width: 1600px) {
  .login-container {
    padding: 4rem;
  }

  .login-card {
    max-width: 550px;
    padding: 4rem;
  }

  .login-header h2 {
    font-size: 2.25rem;
  }

  .login-header p {
    font-size: 1.125rem;
  }

  .input-field {
    font-size: 1.0625rem;
    height: 3.5rem;
  }

  .login-button {
    font-size: 1.125rem;
    height: 3.75rem;
  }
}

/* Ultra-wide desktop enhancements */
@media (min-width: 2000px) {
  .login-container {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0;
    margin: 0;
    background-color: #f5f3ff;
    background-size: cover;
    background-position: center;
  }

  .dark-mode.login-container {
    background-color: #1e1b4b;
  }

  .login-card {
    max-width: 650px;
    width: 100%;
    padding: 5rem;
    margin: 0;
    border-radius: 24px;
    box-shadow: 0 25px 50px rgba(0, 0, 0, 0.15);
  }

  .login-header h2 {
    font-size: 2.5rem;
    margin-bottom: 1.5rem;
  }

  .login-header p {
    font-size: 1.25rem;
    line-height: 1.6;
  }

  .input-group label {
    font-size: 1.125rem;
    margin-bottom: 1rem;
  }

  .input-field {
    padding: 1.25rem 1.5rem 1.25rem 4rem;
    font-size: 1.125rem;
    height: 4rem;
    border-radius: 12px;
  }

  .input-icon {
    left: 1.5rem;
    width: 2rem;
    height: 2rem;
  }

  .toggle-password {
    right: 1.5rem;
    width: 2rem;
    height: 2rem;
  }

  .login-button {
    padding: 1.5rem;
    height: 4.5rem;
    font-size: 1.25rem;
    border-radius: 12px;
    margin-top: 1rem;
  }

  .form-options {
    font-size: 1.125rem;
    margin-top: 0;
  }

  .remember-me input {
    width: 1.25rem;
    height: 1.25rem;
  }

  .error-message {
    font-size: 1.125rem;
    padding: 1rem;
  }

  .login-footer {
    font-size: 1.125rem;
    margin-top: 2.5rem;
  }

  .theme-switcher {
    width: 2.5rem;
    height: 2.5rem;
    top: 2rem;
    right: 2rem;
  }

  .theme-switcher svg {
    width: 2rem;
    height: 2rem;
  }

  .dark-mode .login-card {
    box-shadow: 0 25px 50px rgba(0, 0, 0, 0.4);
    background-color: #312e81;
  }

  .dark-mode .input-field {
    border-width: 2px;
    background-color: #4338ca;
    border-color: #4c1d95;
  }

  .dark-mode .login-button {
    box-shadow: 0 10px 25px rgba(124, 58, 237, 0.3);
  }

  .dark-mode .login-button:hover {
    box-shadow: 0 15px 30px rgba(139, 92, 246, 0.4);
  }
}

@media (min-width: 3000px) {
  .login-card {
    max-width: 700px;
    padding: 6rem;
  }

  .login-header h2 {
    font-size: 3rem;
  }

  .login-header p {
    font-size: 1.5rem;
  }

  .input-field {
    height: 5rem;
    font-size: 1.375rem;
    padding: 1.5rem 2rem 1.5rem 5rem;
  }

  .login-button {
    height: 5.5rem;
    font-size: 1.5rem;
  }

  .theme-switcher {
    width: 3rem;
    height: 3rem;
  }

  .theme-switcher svg {
    width: 2.5rem;
    height: 2.5rem;
  }
}

.login-card {
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.input-field, .login-button {
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}

@media (min-width: 2000px) {
  .dark-mode .login-card {
    box-shadow: 0 25px 50px rgba(0, 0, 0, 0.4);
  }

  .dark-mode .input-field {
    border-width: 2px;
  }

  .dark-mode .login-button {
    box-shadow: 0 10px 25px rgba(124, 58, 237, 0.3);
  }

  .dark-mode .login-button:hover {
    box-shadow: 0 15px 30px rgba(139, 92, 246, 0.4);
  }
}
</style>
