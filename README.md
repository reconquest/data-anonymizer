# data-anonymizer

A dead simple anonymizer for JSON datasets based on a list of rules.

Requires a config in the specified format:
```
key: value
```

For example:
```
contactDetails.company: company
contactDetails.technicalContact.name: name
contactDetails.technicalContact.email: email
contactDetails.technicalContact.city: city
contactDetails.technicalContact.postcode: zip
contactDetails.technicalContact.state: state
addonLicenseId: int_string
licenseId: sen_license
```

List of available generators:
- name
- company
- email
- int_string _generates integer but outputs it as a string_
- city
- zip
- sen_license
- state

# Usage

```bash
data-anonymizer < dataset.json > dataset.anon.json
```

# Installation

Pretty straightforward
```
go get github.com/reconquest/data-anonymizer
```
