import { createLocalVue, shallowMount } from '@vue/test-utils'
import Home from '@/components/Home.vue'
import flushPromises from 'flush-promises'
import VueRouter from 'vue-router'

const localVue = createLocalVue()
localVue.use(VueRouter)
const router = new VueRouter()

describe('Home.vue', () => {
  it('renders home screen', async () => {
    const wrapper = shallowMount(Home, { localVue, router })

    await flushPromises()
    expect(wrapper).toMatchSnapshot()
  })
})
