import { mount } from '@vue/test-utils'
import App from '@/App.vue'
import router from '@/router'

describe('App.vue', () => {
  it('renders main app screen', async () => {
    const wrapper = mount(App, {
      global: {
        plugins: [router]
      }
    })

    expect(wrapper.html()).toMatchSnapshot()
  })

  it('renders navigates to about', async () => {
    const wrapper = mount(App, {
      global: {
        plugins: [router]
      }
    })

    router.push('/about')
    await router.isReady()
    expect(wrapper.html()).toMatchSnapshot()
  })
})
