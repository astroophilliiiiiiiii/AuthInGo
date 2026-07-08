package db

//Repositories box" 📦
// jitne bhi repositories hain unko ek jagah collect karke app ko dena.
// ✅ Taaki app ko database ke saare functions ka access mil jaye ek hi jagah se. ✅
// Agar Storage nahi banaya toh app me alag-alag dena padega:
//  Application{ userRepo,productRepo,orderRepo,paymentRepo } --messy ❌
//  Storage | UserRepo |-- ProductRepo |-- OrderRepo|
// Store: Storage   clean✅

// facilitates dependency injection for repository
type Storage struct {
	UserRespository UserRespository //Storage box ke andar 1 slot hoga jiska naam UserRepository h
}

// constructor for it -- contains objects for all of it
func NewStorage() *Storage {
	return &Storage{
		UserRespository: &UserRespositoryImpl{}, // actual repo object given
	}
}

// App poore project ka main container hai.
//  Uske andar config, database, router sab rakh dete hain taaki sab connected rahe. ✅
// Ab kahin bhi app se repo access ho sakti hai: app.Store.UserRepository.Create()
