# Makefile for generating mocks using mockgen

# Define directories to process
DIRS = controllers repo service

# Target to generate all mocks
.PHONY: mocks
mocks:
	@for %%d in ($(DIRS)) do ( \
		echo Processing %%d directory... && \
		if exist "%%d\mocks" ( \
			echo mocks folder exists in %%d, generating mocks... && \
			for /f "delims=" %%f in ('dir /b /a-d "%%d\*.go" 2^>nul') do ( \
				echo %%~nf | findstr /I "impl" >nul && ( \
					echo Skipping %%d\%%f because filename contains 'impl' \
				) || ( \
					if not "%%~nf" == "mocks" ( \
						echo Generating mock for %%d\%%f && \
						mockgen -source=%%d\%%f -destination=%%d\mocks\mock_%%~nf.go -package=mocks \
					) \
				) \
			) \
		) else ( \
			echo mocks folder does not exist in %%d, creating and generating mocks... && \
			mkdir "%%d\mocks" && \
			for /f "delims=" %%f in ('dir /b /a-d "%%d\*.go" 2^>nul') do ( \
				echo %%~nf | findstr /I "impl" >nul && ( \
					echo Skipping %%d\%%f because filename contains 'impl' \
				) || ( \
					if not "%%~nf" == "mocks" ( \
						echo Generating mock for %%d\%%f && \
						mockgen -source=%%d\%%f -destination=%%d\mocks\mock_%%~nf.go -package=mocks \
					) \
				) \
			) \
		) \
	)

# Clean target to remove all generated mocks
.PHONY: clean-mocks
clean-mocks:
	@for %%d in ($(DIRS)) do ( \
		if exist "%%d\mocks" ( \
			echo Cleaning %%d\mocks directory... && \
			del /Q "%%d\mocks\*.go" 2>nul || echo No mocks to clean in %%d. \
		) \
	)