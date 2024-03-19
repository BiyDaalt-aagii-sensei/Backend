package firebase

type FireStore struct {
	*FireDB
}

func NewStore() *FireStore {
	d := FirebaseDB()
	return &FireStore{
		FireDB: d,
	}
}
