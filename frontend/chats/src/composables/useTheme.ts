import { ref, watch } from 'vue'

const darkMode = ref(true)

export function useTheme() {
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

  function updateTheme() {
    if (darkMode.value) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }

  function toggleTheme() {
    darkMode.value = !darkMode.value
    updateTheme()
    setCookie('theme', darkMode.value ? 'dark' : 'light')
  }

  // Инициализация темы при создании composable
  const themeCookie = getCookie('theme')
  if (themeCookie) {
    darkMode.value = themeCookie === 'dark'
  } else {
    darkMode.value = window.matchMedia('(prefers-color-scheme: dark)').matches
  }
  updateTheme()

  return {
    darkMode,
    toggleTheme
  }
}