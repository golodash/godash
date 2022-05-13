CHANGELOG
=========

UNRELEASED
----------

* ✅ test(slices-tail): tail test cases and benchmarks added
* 🎉 feat(slices-tail): tail function added
* 🐛 fix(slices-same): same function updated so it can compare struct unexported values too
* 🐛 fix(slices-slice): added a new scenario test and satisfied it in slice function
* 🐛 fix(slices-slice): bug fix
* ✅ test(slices-slice_test): 'slice' function's test cases and benchmarks added
* 🎉 feat(slices-slice): 'slice' function added
* ✅ test(slices-reverse): reverse function's test cases and benchmarks added
* 🐛 fix(slices-reverse): reverse function's bug fixed
* 🎉 feat(slices-reverse): reverse function added
* ✅ test(slices-remove): 'remove' function's test cases and benchmarks added
* 🎉 feat(slices-remove): 'remove' function added
* ✅ test(slices-sorted_unique_by): sorted_unique_by test cases and benchmarks added
* 🐛 fix(slices-sorted_unique_by): sorted_unique_by approach changed
* 🎉 feat(slices-sorted_unique_by): 'SortedUniqueBy' function added
* ✅ test(slices-sorted_unique): 'Sorted_Unique' function's test cases and benchmarks added
* 🎉 feat(slices-sorted_unique): 'Sorted_Unique' function added
* 🐛 fix(slices): sorted_last_index_by and sorted_last_index bug fixed
* 🐛 fix(slices-sorted_index_by): bug fix
* 🐛 fix(slices-sorted_index): sorted_index function bug fix
* ✅ test(slices-sorted_last_index_of): sorted_last_index_of test cases and benchmarks added
* 🎉 feat(slices-sorted_last_index_of): sorted_last_index_of function added
* 🐛 fix(slices-sorted_index): sorted_index function bug on behaviour fixed
* ✅ test(slices-sorted_last_index_by): sorted_last_index_by test cases and benchmarks added
* 🎉 feat(slices-sorted_last_index_by): sorted_last_index_by function added
* ✅ test(slices-sorted_last_index): sorted_last_index test cases and benchmarks added
* 🎉 feat(slices-sorted_last_index): sorted_last_index function added
* 🐛 fix(slices): sorted_index* functions meant to return lowest index of passed value but it was returning the highest
* 🐛 fix(slices-sorted_index_of): sorted_index_of behaviour fixed on not found senario
* ✅ test(slices-sorted_index_of): sorted_index_of test cases and benchmarks added
* 🎉 feat(slices-sorted_index_of): sorted_index_of function added
* 🎉 feat(slices-sorted_index_by): sorted_index_by test cases and benchmarks added
* 🎉 feat(internal-utilities): is_number_type function added
* 🎉 feat(slices-sorted_index_by): sorted_index_by function added
* 🐛 fix: doc updated and bug fixed
* 🎉 feat(internal-is_number): is_number function added
* ✅ test(slices-sorted_index): sorted_index test cases and benchmarks added
* 🎉 feat(slices-sorted_index): sorted_index function added
* 🐛 fix(internal): check_same_type function debug
* 🎉 feat(slices-pull_at): pull_at function with test cases and benchmarks added
* 🎉 feat(internal-unique_int): unique_int function added
* 🎉 feat(slices-pull): pull function added
* 🎉 feat(slices-last_index_of): last_index_of function added
* 🎉 feat(slices-last): last function added
* ✅ test(slices-n_th_test): n_th function's benchmark added
* ✅ test(slices-n_th_test): n_th function's test case added
* 🐛 fix(slices-n_th): n_th function's bug fixed on empty slices
* 🎉 feat(slices-n_th): n_th function added
* ✅ test(slices-join_test): join function's test cases and benchmarks added
* 🎉 feat(slices-join): join function added
* 🐛 fix: added fix in git hooks scripts for windows
* ✅ test(slices-intersection_by): intersection_by test cases and benchmarks added
* 🎉 feat(slices-intersection_by): intersection_by function added
* ✅ test(slices-intersection): intersection test cases and benchmarks added
* 🎉 feat(slices-intersection): intersection function added
* ✅ test(slices-from_pairs): from_pairs test cases and benchmarks added
* 🎉 feat(slices-from_pairs): from_pairs function added
* 🎉 feat(slices-same): same function updated, debuged and ignores unexported struct fields
* ✅ test(slices-index_of): index_of test cases and benchmarks added
* 🎉 feat(slices-index_of): index_of function added
* ✅ test(slices-initial): initial test cases and benchmarks added
* 🎉 feat(slices-initial): initial function added
* 🐛 fix: bug fix on test file
* ✅ test(slices-head_first): head and first test cases and benchmarks added
* 🎉 feat(slices-head_first): head and first functions added
* 🐛 fix(slices-latest): a big bug fixed in latest(it was returning the first element)
* ✅ test(slices-flatten): flatten test cases and benchmarks added
* ✅ test(slices-flatten_depth): flatten_depth test cases and benchmarks added
* 🎉 feat(slices-flatten_depth): flatten_depth function added
* ✅ test(slices-flatten_deep): flatten_deep test cases and benchmarks added
* 🎉 feat(slices-flatten_deep): flatten_deep function added
* 🐛 fix(slices-latest): a big bug fixed in latest(it was returning the first element)
* 🎉 feat(slices-flatten): flatten function added
* ✅ test(slices-find_index): find_index test cases and benchmarks added
* 🎉 feat(slices-find_index): find_index function added
* 🐛 fix: fixing the same problem on other test cases
* ✅ test(slices-fill): fill test cases and benchmarks added
* 🎉 feat(slices-fill): fill function added
* 🐛 fix: bug in drop_by_test fixed
* ✅ test(slices-drop_by): drop_by test cases and benchmarks added
* 🎉 feat(slices-drop_by): drop_by function added
* ✅ test(slices-drop_right): drop_right test cases and benchmarks added
* 🎉 feat(slices-drop_right): drop_right function added
* 🐛 fix: bug on test cases fixed
* 🐛 fix: difference and difference_by functions debuged
* ✅ test(slices-drop): made test cases and benchmarks for drop function
* 🎉 feat(slices-drop): drop function added
* ✅ test(slices-difference_by): difference_by test cases and benchmarks added
* 🎉 feat(slices-difference_by): difference_by function added
* 🐛 fix: test cases been running on go 1.18
* 🎉 feat: main workflow for test cases added
* 🎉 feat: husky now removed with a simple script
* ✅ test(slices-difference): difference test cases and benchmarks added
* 🎉 feat(slices-difference): difference function added
* ✅ test(slices-concat): made test cases and benchmarks for concat function
* 🎉 feat(slices-concat): concat function added
* 🐛 fix: git scripts now function properly
* ✅ test(slices-compact): made test cases and benchmarks for compact function
* 🎉 feat(slices-compact): compact function added
* ✅ test(slices-chunk): made test cases and benchmarks more accurate
* 🎉 feat(slices-chunk): chunk function added, updated, and optimized

.. 1.0.0 (yyyy-mm-dd)
.. ------------------
