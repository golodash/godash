CHANGELOG
=========

UNRELEASED
----------

* 🐛 fix(numbers-inrange): correct name
* 🎉 feat(number-inrange): InRange function added
* 🐛 fix(maths): output of Pow in some scenarious updated and testcases updated
* 🐛 fix(maths): math package name updated to maths
* ✅ test(strings-pad_test): Pad testcases and benchmarks added
* 🎉 feat(strings-pad): Pad function added
* ✅ test(strings-pad_end_test): PadEnd testcases and benchmarks added
* 🎉 feat(strings-pad_end): PadEnd function added
* ✅ test(strings-pad_start_test): a new test case added for PadStart function
* ✅ test(strings-pad_start_test): PadStart testcases and benchmarks added
* 🎉 feat(strings-pad_start): PadStart function added
* ✅ test(strings-truncate_test): Truncate testcases and benchamrks added
* 🎉 feat(strings-truncate): Truncate function added
* 🐛 fix(internal-utilities): isCustomSeparator function added to internal package
* ✅ test(strings-words_test): Words testcases and benchmarks added
* 🎉 feat(strings-words): Words function added
* ✅ test(strings-upper_first_test): UpperFirst testcases and benchmarks added
* 🎉 feat(strings-upper_first): UpperFirst function added
* ✅ test(strings-lower_first_test): LowerFirst testcases and benchmarks added
* 🎉 feat(strings-lower_first): LowerFirst function added
* ✅ test(strings-start_case_test): StartCase testcases and benchmarks added
* 🎉 feat(strings-start_case): StartCase function added
* 🚀 perf(strings-repeat): Repeat function optimization happened
* 🚀 perf(strings-lower_case): removed an extra function from code execution
* ✅ test(strings-lower_case_test): LowerCase testcases and benchmarks added
* 🎉 feat(strings-lower_case): LowerCase function added
* ✅ test(strings-repeat): Repeat testcases and benchmarks addeed
* 🎉 feat(strings-repeat): Repeat function added
* ✅ test(strings-starts_with): StartsWith testcases and benchmarks added
* 🎉 feat(strings-starts_with): StartsWith function added
* ✅ test(strings-ends_with): added a new testcase to EndsWith function
* ✅ test(strings-ends_with): EndsWith testcases and benchmarks added
* 🎉 feat(strings-ends_with): EndsWith function added
* ✅ test(strings-pascal_case_test): PascalCase testcases and benchmarks added
* 🎉 feat(strings-pascal_case): PascalCase function added
* ✅ test(strings-camel_case_test): CamelCase testcases and benchmarks added
* 🎉 feat(strings-camel_case): CamelCase function added
* 🎉 feat(strings-utils): internalCamelCase function added for camel and pascal casing solution
* 🎉 feat(strings-utils): updating ScreamingDelimited to CustomDelimitedCase
* ✅ test(strings-kebab_case_test): KebabCase testcases and benchmarks added
* 🎉 feat(strings-kebab_case): KebabCase function added
* ✅ test(strings-snake_case_test): SnakeCase testcases and banchmarks added
* 🎉 feat(strings-snake_case): SnakeCase function added
* 🎉 feat(strings-utils): a general solution function called ToScreamingDelimited added to solve out case changing subject
* ✅ test(slices-sum_by_test): SumBy testcases and benchmarks added
* 🎉 feat(slices-sum_by): SumBy function added
* ✅ test(slices-sum_test): Sum testcases and benchmarks added
* 🎉 feat(slices-sum): Sum function added
* 🐛 fix(slices-mean): Mean and MeanBy functions fixed
* ✅ test(slices-mean_by_test): MeanBy testcases and benchmarks added
* 🎉 feat(slices-mean_by): MeanBy function added
* ✅ test(slices-mean): Mean usescases and benchmarks added
* 🎉 feat(slices-mean): Mean function added
* 🐛 fix: uint types didn't get checked before but now they do get checked
* 🎉 feat(internal): CanUint function added
* ✅ test(maths-multiply): Multiply testcases and benchmarks added
* 🎉 feat(maths-multiply): Multiply function added
* ✅ test(maths-round): Round testcases and benchmarks added
* 🎉 feat(maths-round): Round function added
* ✅ test(maths): added two test cases for Floor and Ceil functions
* 🐛 fix(maths-power): fixed returning float if output was float
* ✅ test(maths-subtract): subtract testcases and benchmarks added
* 🎉 feat(maths-subtract): subtract function added
* ✅ test(slices-min_by): MinBy testcases and benchmarks added
* 🎉 feat(slices-min_by): MinBy function added
* ✅ test(slices-min): Min testcases and benchmarks added
* 🎉 feat(slices-min): Min function added
* ✅ test(slices-max_by): MaxBy testcases and benchmarks added
* 🎉 feat(slices-max_by): MaxBy function added
* ✅ test(slices-max): Max testcases and benchmarks added
* 🎉 feat(slices-max): Max function added
* 🎉 feat(internal): added CompareNumbers and GetNumberTypeRank and GetOutputNumberType functions
* ✅ test(maths-floor): Floor benchmarks and test cases added
* 🎉 feat(maths-floor): Floor function added
* ✅ test(math-divide): divide testcases and benchmarks added
* 🎉 feat(math-divide): divide function added
* ✅ test(math-add): Add testcases and benchmarks added
* 🎉 feat(math-add): Add function added
* 🐛 fix(math): CanInt and CanFloat added to internal package to fix go 1.17 not having those functions inside reflect package
* ✅ test(math-ceil): ceil testcases and benchmarks added
* 🎉 feat(math-ceil): ceil function added
* ✅ test(math-power): power testcases and benchmarks added
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
