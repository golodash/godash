CHANGELOG
=========

UNRELEASED
----------

* 🎉 feat(functions): wrap_func added

1.2.0 (2022-10-26)
------------------

* 🎉 feat(generals-duplicate): Duplicate function added
* 🎉 feat(generals-same): Same function added

1.1.1 (2022-08-13)
------------------

* 🐛 fix(slices-mean_by): MeanBy output expectation updated
* 🐛 fix: bug fix on Mean and MaxBy and some changes on Intersection functions happened

1.1.0 (2022-08-13)
------------------

* 🎉 feat(strings-pad): Pad function added
* ✅ test(slices): all test cases and benchmarks added for numbers, strings and maths packages
* 🎉 feat(strings-pad_end): PadEnd function added
* 🎉 feat(strings-pad_start): PadStart function added
* 🎉 feat(strings-truncate): Truncate function added
* 🎉 feat(strings-words): Words function added
* 🎉 feat(strings-upper_first): UpperFirst function added
* 🎉 feat(strings-lower_first): LowerFirst function added
* 🎉 feat(strings-start_case): StartCase function added
* 🚀 perf(strings-repeat): Repeat function optimization happened
* 🚀 perf(strings-lower_case): removed an extra function from code execution
* 🎉 feat(strings-lower_case): LowerCase function added
* 🎉 feat(strings-repeat): Repeat function added
* 🎉 feat(strings-starts_with): StartsWith function added
* 🎉 feat(strings-ends_with): EndsWith function added
* 🎉 feat(strings-pascal_case): PascalCase function added
* 🎉 feat(strings-camel_case): CamelCase function added
* 🎉 feat(strings-utils): internalCamelCase function added for camel and pascal casing solution
* 🎉 feat(strings-utils): updating ScreamingDelimited to CustomDelimitedCase
* 🎉 feat(strings-kebab_case): KebabCase function added
* 🎉 feat(strings-snake_case): SnakeCase function added
* 🎉 feat(strings-utils): a general solution function called ToScreamingDelimited added to solve out case changing subject
* 🎉 feat(slices-sum_by): SumBy function added
* 🎉 feat(slices-sum): Sum function added
* 🎉 feat(slices-mean_by): MeanBy function added
* 🎉 feat(slices-mean): Mean function added
* 🎉 feat(internal): CanUint function added
* 🎉 feat(maths-multiply): Multiply function added
* 🎉 feat(maths-round): Round function added
* 🎉 feat(maths-subtract): subtract function added
* 🎉 feat(slices-min_by): MinBy function added
* 🎉 feat(slices-min): Min function added
* 🎉 feat(slices-max_by): MaxBy function added
* 🎉 feat(slices-max): Max function added
* 🎉 feat(internal): added CompareNumbers and GetNumberTypeRank and GetOutputNumberType functions
* 🎉 feat(maths-floor): Floor function added
* 🎉 feat(math-divide): divide function added
* 🎉 feat(math-add): Add function added
* 🎉 feat(internal): CanInt and CanFloat added to internal package to fix go 1.17 not having those functions inside reflect package
* 🎉 feat(math-ceil): ceil function added
* 🎉 feat(math-power): power function added

1.0.1 (2022-07-19)
------------------

* 🐛 fix(slices): bug on test cases, documentations and some functions resolved

1.0.0 (2022-06-22)
------------------

* 🎉 feat(internal): internal same function added to compare two different variable without breaking the code
* 🎉 feat(internal): some other functions added: 1.slice_check 2.check_same_type 3.are_comparable 4.unique_int 5.is_number 6.is_number_type
* ✅ test(slices): all test cases and benchmarks added for slices package
* 🎉 feat(slices-find_index_by): find_index_by function added
* 🎉 feat(slices-take_while): take_while function added
* 🎉 feat(slices-take_right_while): take_right_while function added
* 🎉 feat(slices-union_by): union_by function added
* 🎉 feat(slices-union): 'union' function added
* 🎉 feat(slices-xor_by): xor_by function added
* 🎉 feat(slices-xor): xor new function added
* 🎉 feat(slices-zip_by): zip_by function added
* 🎉 feat(slices-zip_map_deep): zip_map_deep function added
* 🎉 feat(slices-zip_map): zip_map function added
* 🎉 feat(slices-zip): zip function added
* 🎉 feat(slices-unzip): unzip function added
* 🎉 feat(slices-take_right): 'take_right' function added
* 🎉 feat(slices-take): 'take' function added
* 🎉 feat(slices-without): without alias of difference added
* 🎉 feat(slices-unique_by): unique_by function added
* 🎉 feat(slices-unique): unique function added
* 🎉 feat(slices-tail): tail function added
* 🎉 feat(slices-slice): 'slice' function added
* 🎉 feat(slices-reverse): reverse function added
* 🎉 feat(slices-remove_by): 'remove_by' function added
* 🎉 feat(slices-sorted_unique_by): 'SortedUniqueBy' function added
* 🎉 feat(slices-sorted_unique): 'Sorted_Unique' function added
* 🎉 feat(slices-sorted_last_index_of): sorted_last_index_of function added
* 🎉 feat(slices-sorted_last_index_by): sorted_last_index_by function added
* 🎉 feat(slices-sorted_last_index): sorted_last_index function added
* 🎉 feat(slices-sorted_index_of): sorted_index_of function added
* 🎉 feat(slices-sorted_index_by): sorted_index_by test cases and benchmarks added
* 🎉 feat(slices-sorted_index_by): sorted_index_by function added
* 🎉 feat(slices-sorted_index): sorted_index function added
* 🎉 feat(slices-pull_at): pull_at function with test cases and benchmarks added
* 🎉 feat(slices-pull): pull function added
* 🎉 feat(slices-last_index_of): last_index_of function added
* 🎉 feat(slices-last): last function added
* 🎉 feat(slices-n_th): n_th function added
* 🎉 feat(slices-join): join function added
* 🎉 feat(slices-intersection_by): intersection_by function added
* 🎉 feat(slices-intersection): intersection function added
* 🎉 feat(slices-from_pairs): from_pairs function added
* 🎉 feat(slices-same): same function updated, debuged and ignores unexported struct fields
* 🎉 feat(slices-index_of): index_of function added
* 🎉 feat(slices-initial): initial function added
* 🎉 feat(slices-head_first): head and first functions added
* 🎉 feat(slices-flatten_depth): flatten_depth function added
* 🎉 feat(slices-flatten_deep): flatten_deep function added
* 🎉 feat(slices-flatten): flatten function added
* 🎉 feat(slices-find_index): find_index function added
* 🎉 feat(slices-fill): fill function added
* 🎉 feat(slices-drop_by): drop_by function added
* 🎉 feat(slices-drop_right): drop_right function added
* 🎉 feat(slices-drop): drop function added
* 🎉 feat(slices-difference_by): difference_by function added
* 🎉 feat(slices-difference): difference function added
* 🎉 feat(slices-concat): concat function added
* 🎉 feat(slices-compact): compact function added
* 🎉 feat(slices-chunk): chunk function added, updated, and optimized
