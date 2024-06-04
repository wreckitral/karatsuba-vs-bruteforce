benchmark:
	@mkdir -p ../graphic ; \
	echo "Changing directory to fibo" ; \
	cd karatsuba && \
	echo "Running Go benchmarks" && \
	go test -bench=. | tee ../graphic/out.dat && \
	echo "Benchmark results:" && \
	cat ../graphic/out.dat && \
	echo "Processing benchmark output" && \
	awk '/Benchmark/{count ++; funcname=$$1; gsub(/Benchmark/,"",funcname); printf("%d,%s,%s,%s,%s\n",count,funcname,$$2,$$3,$$4)}' ../graphic/out.dat > ../graphic/final.dat && \
	echo "Processed benchmark data:" && \
	cat ../graphic/final.dat && \
	echo "Generating operations graphic" && \
	gnuplot -e "file_path='../graphic/final.dat'" \
	        -e "graphic_file_name='../graphic/operations.png'" \
	        -e "y_label='number of operations'" \
	        -e "y_range_min=0" \
	        -e "y_range_max=30000000" \
	        -e "column_1=1" \
	        -e "column_2=3" \
	        -e "column_3=2" ../graphic/performance.gp && \
	echo "Generating time operations graphic" && \
	gnuplot -e "file_path='../graphic/final.dat'" \
	        -e "graphic_file_name='../graphic/time_operations.png'" \
	        -e "y_label='each operation in nanoseconds'" \
	        -e "y_range_min=0" \
	        -e "y_range_max=3500" \
	        -e "column_1=1" \
	        -e "column_2=4" \
	        -e "column_3=2" ../graphic/performance.gp && \
	echo "Cleaning up" && \
	rm -f ../graphic/out.dat ../graphic/final.dat && \
	echo "'graphic/operations.png' and 'graphic/time_operations.png' graphics were generated."
