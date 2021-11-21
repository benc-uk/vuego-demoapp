import { mount } from '@vue/test-utils'
import Home from '@/components/Home.vue'

describe('Home.vue', () => {
  it('renders home screen', async () => {
    const wrapper = mount(Home, {})

    expect(wrapper.html()).toMatchSnapshot()
  })
})
