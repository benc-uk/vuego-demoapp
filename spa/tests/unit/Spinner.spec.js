import { shallowMount } from '@vue/test-utils'
import Spinner from '@/components/Spinner.vue'
import flushPromises from 'flush-promises'

describe('Spinner.vue', () => {
  it('renders a spinner', async () => {
    const wrapper = shallowMount(Spinner)

    await flushPromises()
    expect(wrapper).toMatchSnapshot()
  })
})
