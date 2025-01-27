
import { registerPlugins } from '@/plugins'
import App from './App.vue'
import { createApp } from 'vue'
import router from './router'

const app = createApp(App)

registerPlugins(app)

// app.mount('#app')
router.isReady().then(() => {
    app.mount('#app')
})
