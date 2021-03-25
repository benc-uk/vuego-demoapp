import { createLocalVue, shallowMount } from '@vue/test-utils'
import App from '@/App.vue'
import flushPromises from 'flush-promises'
import VueRouter from 'vue-router'
//import { toHaveRouteName } from 'vue-test-utils-helpers'

const localVue = createLocalVue()
localVue.use(VueRouter)
const router = new VueRouter()

describe('App.vue', () => {
  it('renders main app screen', async () => {
    const wrapper = shallowMount(App, { localVue, router })

    wrapper.find('b-button[to="/home"]').trigger('click')
    await wrapper.vm.$nextTick()
    await flushPromises()

    //console.log(wrapper.vm.$route)
    //expect(wrapper.vm.$route.name).toBe('home')

    await flushPromises()
    expect(wrapper).toMatchSnapshot()
  })
})
