import { createLocalVue, shallowMount } from '@vue/test-utils'
import About from '@/components/About.vue'
import flushPromises from 'flush-promises'

const localVue = createLocalVue()

describe('About.vue', () => {
  it('renders about screen', async () => {
    const wrapper = shallowMount(About, { localVue })

    await flushPromises()
    expect(wrapper).toMatchSnapshot()
  })
})
