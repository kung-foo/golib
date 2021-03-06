# Tideland Go Library

## 2015-08-18

- Monitoring in cells is now pluggable

## 2015-08-17

- Fixed race condition in cells
- Optimised time handling in cells

## 2015-08-09

- Added `Collect()` and `DoAll()` to errors

## 2015-08-02

- Added `BeginOf()` and `EndOf()` to timex

## 2015-08-01

- Added `Set` and `StringSet` to collections
- Added `Retry()` to timex

## 2015-07-28

- Added assertion `Retry()` to audit

## 2015-07-26

- Added `CallbackBehavior` to cells

## 2015-07-24

- Fixed cells unsubscribing failure when stopping cell with 
  bi-directional subscriptions; thanks to Jonathan Camp for
  his fix
- Added expected value to compare with signal in `Wait()` assertion
- Added test for configuration validation in configurator behavior

## 2015-07-23

- Added `ReadFile()` to configuration
- Added `SimpleProcessorBehavior` to cells 
- Added `ConfiguratorBehavior` to cells
- Added assertion `Wait()` to audit

## 2015-07-20

- Simplified `configuration` package for usage with `stringex.Defaulter`

## 2015-07-17

- Added `stringex` package

## 2015-07-10

- Added `KeyStringValueTreeBuilder` to SML
- Several minor fixes

## 2015-07-05

- Made standard logger backend time format changeable

## 2015-06-28

- Changed configuration API to be more powerful
  and convenient

## 2015-06-25

- Added new `SceneBehavior` to cells

## 2015-06-25

- Done migration into new library
- Added new configuration package

## 2015-06-05

- Started migration of existing packages into new library
