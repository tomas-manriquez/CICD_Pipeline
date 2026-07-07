module cicd_pipeline/backend

go 1.26.4

replace cicd_pipeline/utils => ./utils

require (
	cicd_pipeline/stages v0.0.0-00010101000000-000000000000
	cicd_pipeline/utils v0.0.0-00010101000000-000000000000
)

replace cicd_pipeline/stages => ./stages
