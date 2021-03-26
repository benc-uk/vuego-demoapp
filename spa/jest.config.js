module.exports = {
  preset: '@vue/cli-plugin-unit-jest/presets/no-babel',
  setupFiles: ['<rootDir>/tests/unit/jest.init.js', 'jest-canvas-mock'],
  verbose: true,
  silent: true,

  collectCoverage: true,
  coverageReporters: ['html', 'text-summary'],

  reporters: [
    'default',
    [
      './node_modules/jest-html-reporter',
      {
        pageTitle: 'Test Report'
      }
    ]
  ],

  collectCoverageFrom: ['src/**/*.vue']
}
