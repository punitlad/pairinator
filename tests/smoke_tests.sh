#!/usr/bin/env bats

@test "validate pairinator rotates" {
  run bash -c "./pairinator add Bender"
  run bash -c "./pairinator add Leela"
  run bash -c "./pairinator add Zoidberg"
  run bash -c "./pairinator add Fry"

  run bash -c "./pairinator rotate"
  [[ "${status}" == 0 ]]
  [[ "${output}" =~ "Good news everyone! Pairs have been rotated!" ]]
}

teardown() {
  rm pairinator.db
}
