module.exports = {
  preset: '@vue/cli-plugin-unit-jest/presets/no-babel',
  setupFiles: ['jest-canvas-mock'],
  verbose: true,
  silent: true,
  collectCoverage: true,
  coverageReporters: ['html', 'text-summary'],
  transform: {
    '^.+\\.vue$': 'vue-jest'
  },
  moduleNameMapper: {
    '\\.(css|less|svg)$': '<rootDir>/tests/unit/__mocks__/styleMock.js'
  }
}
