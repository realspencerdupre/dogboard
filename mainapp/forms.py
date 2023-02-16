from django import forms

from .models import Sound


class SoundForm(forms.ModelForm):
    audio = forms.FileField(widget=forms.ClearableFileInput)
    icon = forms.ImageField(widget=forms.ClearableFileInput)

    class Meta:
        model = Sound
        fields = ["name", "audio", "icon", "description", "source"]
        widgets = {
            'name': forms.TextInput(attrs={'class': 'form-control'}),
            'audio': forms.ClearableFileInput(attrs={'class': 'form-control'}),
            'icon': forms.ClearableFileInput(attrs={'class': 'form-control'}),
            'description': forms.Textarea(attrs={'class': 'form-control', 'rows': 3}),
            'source': forms.TextInput(attrs={'class': 'form-control'}),
        }
