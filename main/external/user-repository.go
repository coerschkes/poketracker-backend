package external

const (
	selectUserQuery = "SELECT id FROM useraccount WHERE firebaseuid = $1"
)

type UserRepository interface {
	Find(firebaseUid string) (int, error)
	Create(firebaseUid string, mail string) error
}

type UserRepositoryImpl struct {
	connector *DatabaseConnector
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{connector: NewDatabaseConnector()}
}

func (i *UserRepositoryImpl) Find(firebaseUid string) (int, error) {
	query, err := i.connector.Query(selectUserQuery, NewUserMapper(), firebaseUid)
	if err != nil {
		return 0, err
	}
	return query.(int), nil
}

func (i *UserRepositoryImpl) Create(firebaseUid string, mail string) error {
	_, err := i.Find(firebaseUid)
	if err != nil {
		_, err := i.connector.Execute("INSERT INTO useraccount (firebaseuid, mail) VALUES ($1, $2)", firebaseUid, mail)
		if err != nil {
			return err
		}
	}
	return err
}
