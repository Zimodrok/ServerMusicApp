import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router'
import './style.css'

async function bootstrap() {
  const envBase = (import.meta as any).env?.VITE_API_BASE?.trim?.() || ''
  let apiBase = envBase
  let cfg: any = null

  if (!apiBase) {
    const host = window.location.hostname || 'localhost'
    const protocol = window.location.protocol || 'http:'
    const currentPort = window.location.port
    const candidates: string[] = []

    // Try API port from heuristic (frontendPort - 1) first.
    if (currentPort) {
      const p = parseInt(currentPort, 10)
      if (!isNaN(p) && p > 0) {
        candidates.push(`${protocol}//${host}:${p - 1}`)
      }
    }
    // Fall back to common default.
    candidates.push(`${protocol}//${host}:8080`)

    for (const base of candidates) {
      try {
        const res = await fetch(`${base}/config/ports`, {
          credentials: 'include',
        })
        if (res.ok) {
          cfg = await res.json()
          apiBase = `${protocol}//${host}:${cfg.api_port || new URL(base).port}`
          break
        }
      } catch (e) {
        // try next candidate
      }
    }
  }

  ;(window as any).__MUSICAPP_API_BASE = apiBase || undefined
  ;(window as any).__MUSICAPP_CONFIG = cfg

  createApp(App).use(router).mount('#app')
}

bootstrap()
