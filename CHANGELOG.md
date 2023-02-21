# Changelog
API Gateway

## [Unreleased]

## [1.7.3] - 20/02/2023
### Fixed
- CORS Allow Method and max Age

## [1.7.2] - 20/02/2023
### Fixed
- CORS Allow Origin

## [1.7.1] - 18/02/2023
### Fixed
- Update Profile service version

## [1.7.0] - 15/02/2023
### Added
- Integration with authorization service (notification)
### Fixed
- Migrate postgresql client to Commons

## [1.6.2] - 20/01/2023
### Fixed
- Remove Deploy Stage
- Change ReplaceSecrets.java to replace-variables.py

## [1.6.1] - 14/01/2023
### Fixed
- Prometheus metric package
- Total request by path metric is changed

## [1.6.0] - 12/01/2023
### Added
- Prometheus metrics

## [1.5.0] - 04/01/2023
### Added
- Integration with profile service (List, Save and Update)

## [1.4.1] - 04/01/2023
### Fixed
- Integration with authorization service with generic handler (all methods)

## [1.4.0] - 03/01/2023
### Added
- Integration with authorization service (List and Update)

## [1.3.0] - 16/11/2022
### Added
- Integration with authorization service (Login, Register and Validate)

## [1.2.1] - 14/10/2022
### Fixed
- Block same version deploy in jenkins

## [1.2.0] - 12/10/2022
### Added
- Routes administration (disabled without auth service integration)

## [1.1.2] - 25/09/2022
### Fixed
- Refactor Jenkinsfile

## [1.1.1] - 20/09/2022
### Fixed
- Clean code

## [1.1.0] - 19/09/2022
### Added
- HTTP Router to any

## [1.0.1] - 28/07/2022
### Fix
- Pipeline version

## [1.0.0] - 14/07/2022
### Added
- Dynamic routing from database (http)
- Reload routing configuration
