
$MyFileName = "\uploads\wallpaper.png"
$WorkingDir = Convert-Path .

$FullPath = $WorkingDir+$MyFileName

Set-ItemProperty -path 'HKCU:\Control Panel\Desktop\' -name wallpaper -value $FullPath

RUNDLL32.EXE user32.dll,UpdatePerUserSystemParameters