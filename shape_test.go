package golang_united_school_homework

import "testing"

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Height: 4, Weight: 5}
	want := float64(20)
	got := r.CalcArea()
	if want != got {
		t.Errorf("error in TestRectangleArea() want: %f got %f", want, got)
	}

	want = float64(18)
	got = r.CalcPerimeter()
	if want != got {
		t.Errorf("error in TestRectangleArea() want: %f got %f", want, got)
	}
}

func TestAddShape(t *testing.T) {
	box := NewBox(1)
	r := Rectangle{Height: 4, Weight: 5}
	err := box.AddShape(r)
	if err != nil {
		t.Errorf("error in TestAddShape() want: %v got %v", nil, err)
	}
	err = box.AddShape(r)
	if err == nil {
		t.Errorf("error in TestAddShape() want: %v got %v", err, nil)
	}
}

func TestGetByIndex(t *testing.T) {
	box := NewBox(1)
	r := Rectangle{Height: 4, Weight: 5}
	box.AddShape(r)
	_, err := box.GetByIndex(0)
	if err != nil {
		t.Errorf("error in TestAddShape() want: %v got %v", nil, err)
	}
	_, err = box.GetByIndex(1)
	if err == nil {
		t.Errorf("error in TestAddShape() want: %v got %v", err, nil)
	}
}
