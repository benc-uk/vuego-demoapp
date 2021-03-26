import { createLocalVue, shallowMount } from '@vue/test-utils'
import Dial from '@/components/Dial.vue'
import flushPromises from 'flush-promises'

const localVue = createLocalVue()

describe('Dial.vue', () => {
  it('renders value on dial', async () => {
    const wrapper = shallowMount(Dial, {
      localVue,
      propsData: {
        value: 22,
        title: 'Test Dial'
      }
    })
    await flushPromises()
    expect(wrapper).toMatchSnapshot()
  })
})
