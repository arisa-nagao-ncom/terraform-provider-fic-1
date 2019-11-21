/*
Package clientconfig provides convenient functions for creating
Flexible InterConnect clients.
It is based on the Python os-client-config library.

Example to Create a Provider Client From clouds.yaml

	opts := &clientconfig.ClientOpts{
		Name: "hawaii",
	}

	pClient, err := clientconfig.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}


Example to Manually Create a Provider Client

	opts := &clientconfig.ClientOpts{
		AuthInfo: &clientconfig.AuthInfo{
			AuthURL:     "https://hi.example.com:5000/v3",
			Username:    "jdoe",
			Password:    "password",
			ProjectName: "Some Project",
			DomainName:  "default",
		},
	}

	pClient, err := clientconfig.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}

*/
package clientconfig
