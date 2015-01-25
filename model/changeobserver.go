package model

type changeObserver interface {
    personAdded(emailName string)
}
