import { mount } from '@vue/test-utils'
import Spinner from '@/components/Spinner.vue'

describe('Spinner.vue', () => {
  it('renders a spinner', async () => {
    const wrapper = mount(Spinner, {})

    expect(wrapper.html()).toMatchSnapshot()
  })
})
