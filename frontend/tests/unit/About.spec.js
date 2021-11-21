import { mount } from '@vue/test-utils'
import About from '@/components/About.vue'

describe('About.vue', () => {
  it('renders about screen', async () => {
    const wrapper = mount(About, {})

    expect(wrapper.html()).toMatchSnapshot()
  })
})
